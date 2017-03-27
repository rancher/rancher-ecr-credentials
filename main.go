package main

import (
	"encoding/base64"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/rancher/go-rancher/client"
)

// Rancher holds the configuration parameters
type Rancher struct {
	URL         string
	AccessKey   string
	SecretKey   string
	RegistryIds []string
	AutoCreate  bool
	client      *client.RancherClient
}

func initLogger() {
	// check if config param has been set for log level, otherwise the default of the logrus package will be used
	if logLevel, ok := os.LookupEnv("LOG_LEVEL"); ok && logLevel != "" {
		logLevelObj, err := log.ParseLevel(logLevel)
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(logLevelObj)
	}
	// set log format to JSON
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	initLogger()
	log.Info("Starting ECR Credential Updater")
	r := Rancher{
		URL:         os.Getenv("CATTLE_URL"),
		AccessKey:   os.Getenv("CATTLE_ACCESS_KEY"),
		SecretKey:   os.Getenv("CATTLE_SECRET_KEY"),
		RegistryIds: []string{},
	}
	if val, ok := os.LookupEnv("AUTO_CREATE"); ok {
		b, err := strconv.ParseBool(val)
		if err != nil {
			log.Fatalf("Unable to parse boolean value from AUTO_CREATE: %s\n", err)
		}
		r.AutoCreate = b
	}
	rancher, err := client.NewRancherClient(&client.ClientOpts{
		Url:       r.URL,
		AccessKey: r.AccessKey,
		SecretKey: r.SecretKey,
	})
	if err != nil {
		log.Fatalf("Unable to create Rancher API client: %s\n", err)
	}
	r.client = rancher
	log.Debug("Created Rancher API Client")

	if ids, ok := os.LookupEnv("AWS_ECR_REGISTRY_IDS"); ok && ids != "" {
		log.Debug("Detected AWS_ECR_REGISTRY_IDS config param")
		r.RegistryIds = strings.Split(ids, ",")
	}

	go healthcheck()

	r.updateEcr(awsClient(), r.client.Registry, r.client.RegistryCredential)
	ticker := time.NewTicker(6 * time.Hour)
	for {
		log.Debug("Sleeping until next poll cycle")
		<-ticker.C
		r.updateEcr(awsClient(), r.client.Registry, r.client.RegistryCredential)
	}
}

func (r *Rancher) updateEcr(
	svc ecriface.ECRAPI,
	registryClient client.RegistryOperations,
	registryCredentialClient client.RegistryCredentialOperations) {

	log.Println("Updating ECR Credentials")

	request := &ecr.GetAuthorizationTokenInput{}
	if len(r.RegistryIds) > 0 {
		request = &ecr.GetAuthorizationTokenInput{RegistryIds: aws.StringSlice(r.RegistryIds)}
	}
	resp, err := svc.GetAuthorizationToken(request)
	log.Debug(resp)
	if err != nil {
		log.Printf("Error calling AWS API: %s\n", err)
		return
	}
	log.Println("Returned from AWS GetAuthorizationToken call successfully")

	if len(resp.AuthorizationData) < 1 {
		log.Println("Request did not return authorization data")
		return
	}

	for _, data := range resp.AuthorizationData {
		r.processToken(data, registryClient, registryCredentialClient)
	}
}

func (r *Rancher) processToken(
	data *ecr.AuthorizationData,
	registryClient client.RegistryOperations,
	registryCredentialClient client.RegistryCredentialOperations) {

	bytes, err := base64.StdEncoding.DecodeString(*data.AuthorizationToken)
	if err != nil {
		log.Printf("[%s] Error decoding authorization token: %s\n", *data.ProxyEndpoint, err)
		return
	}
	token := string(bytes[:len(bytes)])

	authTokens := strings.Split(token, ":")
	if len(authTokens) != 2 {
		log.Printf("[%s] Authorization token does not contain data in <user>:<password> format: %s\n", *data.ProxyEndpoint, token)
		return
	}

	registryURL, err := url.Parse(*data.ProxyEndpoint)
	if err != nil {
		log.Printf("[%s] Error parsing registry URL: %s\n", *data.ProxyEndpoint, err)
		return
	}

	ecrUsername := authTokens[0]
	ecrPassword := authTokens[1]
	ecrHost := registryURL.Host

	registries, err := registryClient.List(&client.ListOpts{})
	if err != nil {
		log.Printf("[%s] Failed to retrieve registries: %s\n", *data.ProxyEndpoint, err)
		return
	}
	log.Printf("[%s] Looking for configured registry for host: %s\n", *data.ProxyEndpoint, ecrHost)
	for _, registry := range registries.Data {
		serverAddress, err := url.Parse(registry.ServerAddress)
		if err != nil {
			log.Printf("[%s] Failed to parse configured registry URL: %s\n", *data.ProxyEndpoint, registry.ServerAddress)
			break
		}
		registryHost := serverAddress.Host
		if registryHost == "" {
			registryHost = serverAddress.Path
		}
		if registryHost == ecrHost {
			credentials, err := registryCredentialClient.List(&client.ListOpts{
				Filters: map[string]interface{}{
					"registryId": registry.Id,
				},
			})
			if err != nil {
				log.Printf("[%s] Failed to retrieved registry credentials for id: %s, %s\n", *data.ProxyEndpoint, registry.Id, err)
				break
			}
			if len(credentials.Data) != 1 {
				log.Printf("[%s] No credentials retrieved for registry: %s\n", *data.ProxyEndpoint, registry.Id)
				break
			}
			credential := credentials.Data[0]
			_, err = registryCredentialClient.Update(&credential, &client.RegistryCredential{
				PublicValue: ecrUsername,
				SecretValue: ecrPassword,
				Email:       "not-really@required.anymore",
			})
			if err != nil {
				log.Printf("[%s] Failed to update registry credential %s, %s\n", *data.ProxyEndpoint, credential.Id, err)
			} else {
				log.Printf("[%s] Successfully updated credentials %s for registry %s; registry address: %s\n", *data.ProxyEndpoint, credential.Id, registry.Id, registryHost)
			}
			return
		}
	}
	log.Printf("[%s] Did not find an existing reigstry for host: %s\n", *data.ProxyEndpoint, ecrHost)

	// If we made it this far, it means we were not able to find an existing registry to update in Rancher
	if r.AutoCreate {
		log.Printf("[%s] Automatically creating registry for host: %s\n", *data.ProxyEndpoint, ecrHost)
		registry, err := registryClient.Create(&client.Registry{
			ServerAddress: ecrHost,
		})
		if err != nil {
			log.Printf("[%s] Error creating registry for host: %s, %s\n", *data.ProxyEndpoint, ecrHost, err)
			return
		}
		_, err = registryCredentialClient.Create(&client.RegistryCredential{
			RegistryId:  registry.Id,
			PublicValue: ecrUsername,
			SecretValue: ecrPassword,
			Email:       "not-really@required.anymore",
		})
		log.Printf("[%s] Successfully created regristy %s and updated credential\n", *data.ProxyEndpoint, registry.Id)

		if err != nil {
			log.Printf("[%s] Error creating registry credential for host: %s, %s\n", *data.ProxyEndpoint, ecrHost, err)
			return
		}
	} else {
		log.Printf("[%s] Failed to find Rancher registry to update for ECR Host: %s\n", *data.ProxyEndpoint, ecrHost)
	}
	return
}

func healthcheck() {

	listenPort := "8080"
	p, ok := os.LookupEnv("LISTEN_PORT")
	if ok {
		listenPort = p
	}
	http.HandleFunc("/ping", ping)
	log.Printf("Starting Healthcheck listener at :%s/ping\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil)
	if err != nil {
		log.Fatal("Error creating health check listener: ", err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Debug("Recieved Health Check Request")
	fmt.Fprintf(w, "pong!")
}

func awsClient() *ecr.ECR {
	roleArn, ok := os.LookupEnv("AWS_ROLE_ARN")
	if ok {
		log.Printf("[awsClient] Assuming Role: %s\n", roleArn)
		return ecr.New(
			session.New(
				aws.NewConfig().WithCredentials(
					stscreds.NewCredentials(session.New(), roleArn),
				),
			),
		)
	}
	return ecr.New(session.New())
}
