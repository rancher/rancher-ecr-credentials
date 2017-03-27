package main

import (
	"encoding/base64"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/rancher/go-rancher/client"
	"github.com/rancher/rancher-ecr-credentials/mocks"
)

func TestMain_basic(t *testing.T) {
	r := &Rancher{}
	mockEcr := new(mocks.ECRAPI)
	mockRegistry := new(mocks.RegistryOperations)
	mockRegistryCredential := new(mocks.RegistryCredentialOperations)
	mockEcr.On("GetAuthorizationToken", &ecr.GetAuthorizationTokenInput{}).Return(
		&ecr.GetAuthorizationTokenOutput{
			AuthorizationData: []*ecr.AuthorizationData{
				&ecr.AuthorizationData{
					ProxyEndpoint:      aws.String("https://012345678910.dkr.ecr.us-east-1.amazonaws.com"),
					AuthorizationToken: aws.String(base64.StdEncoding.EncodeToString([]byte("mockUser:mockPassword"))),
				},
			},
		}, nil)

	mockRegistry.On("List", &client.ListOpts{}).Return(
		&client.RegistryCollection{
			Data: []client.Registry{
				client.Registry{
					Resource: client.Resource{
						Id: "1r1",
					},
					ServerAddress: "012345678910.dkr.ecr.us-east-1.amazonaws.com",
				},
			},
		},
		nil,
	)

	credential := client.RegistryCredential{
		Resource:   client.Resource{Id: "1rc1"},
		RegistryId: "1r1",
	}
	mockRegistryCredential.On("List", &client.ListOpts{
		Filters: map[string]interface{}{
			"registryId": "1r1",
		},
	}).Return(&client.RegistryCredentialCollection{
		Data: []client.RegistryCredential{credential},
	}, nil)
	mockRegistryCredential.On("Update", &credential, &client.RegistryCredential{
		PublicValue: "mockUser",
		SecretValue: "mockPassword",
		Email:       "not-really@required.anymore",
	}).Return(&client.RegistryCredential{}, nil)

	r.updateEcr(mockEcr, mockRegistry, mockRegistryCredential)

	mockEcr.AssertExpectations(t)
	mockRegistry.AssertExpectations(t)
	mockRegistryCredential.AssertExpectations(t)
}

func TestMain_autoCreate(t *testing.T) {
	r := &Rancher{
		AutoCreate: true,
	}
	mockEcr := new(mocks.ECRAPI)
	mockRegistry := new(mocks.RegistryOperations)
	mockRegistryCredential := new(mocks.RegistryCredentialOperations)
	mockEcr.On("GetAuthorizationToken", &ecr.GetAuthorizationTokenInput{}).Return(
		&ecr.GetAuthorizationTokenOutput{
			AuthorizationData: []*ecr.AuthorizationData{
				&ecr.AuthorizationData{
					ProxyEndpoint:      aws.String("https://012345678910.dkr.ecr.us-east-1.amazonaws.com"),
					AuthorizationToken: aws.String(base64.StdEncoding.EncodeToString([]byte("mockUser:mockPassword"))),
				},
			},
		}, nil)

	mockRegistry.On("List", &client.ListOpts{}).Return(
		&client.RegistryCollection{
			Data: []client.Registry{},
		},
		nil,
	)

	mockRegistry.On("Create",
		&client.Registry{
			ServerAddress: "012345678910.dkr.ecr.us-east-1.amazonaws.com",
		},
	).Return(&client.Registry{
		Resource: client.Resource{
			Id: "1r1",
		},
		ServerAddress: "012345678910.dkr.ecr.us-east-1.amazonaws.com",
	}, nil)

	mockRegistryCredential.On("Create", &client.RegistryCredential{
		RegistryId:  "1r1",
		PublicValue: "mockUser",
		SecretValue: "mockPassword",
		Email:       "not-really@required.anymore",
	}).Return(&client.RegistryCredential{
		Resource:    client.Resource{Id: "1rc1"},
		RegistryId:  "1r1",
		PublicValue: "mockUser",
		SecretValue: "mockPassword",
		Email:       "not-really@required.anymore",
	}, nil)

	r.updateEcr(mockEcr, mockRegistry, mockRegistryCredential)

	mockEcr.AssertExpectations(t)
	mockRegistry.AssertExpectations(t)
	mockRegistryCredential.AssertExpectations(t)
}
