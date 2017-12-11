# Rancher ECR Credentials Updater

[![Apache License 2.0](https://img.shields.io/badge/license-Apache_License_2.0-blue.svg)](https://github.com/rancher/rancher-ecr-credentials/blob/master/LICENSE)

This is a Docker container that when executed will update the Docker registry
credentials in Rancher for an Amazon Elastic Container Registry.

Originally contributed by John Engelman from [Object Partners](http://www.objectpartners.com). 

## Why is this needed?

Because access to ECR is controlled with AWS IAM.
An IAM user must request a temporary credential to the registry using the AWS API.
This temporary credential is then valid for 12 hours.

Rancher only supports registries that authenticate with a username and password.

## How to use

In order to authenticate with AWS ECR, this Docker container uses the default
chain of [credential providers](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#config-settings-and-precedence).

The only requirement for running this application is to specify the AWS region
using the `AWS_REGION` environment variable.

AWS credentials are loaded using the default [AWS credential chain](http://docs.aws.amazon.com/sdk-for-go/latest/v1/developerguide/configuring-sdk.title.html).
Credentials are loaded in the following order:

1. Assumed IAM Role specified in `AWS_ROLE_ARN` (The credentials used to execute the assume are determined using the following rules)
1. Environment variables (Specify `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION`, and `AWS_SESSION_TOKEN` *(optional)*)
1. Shared credentials file (mount a volume to `/root/.aws` that contains `credentials` and `config` files and specify `AWS_PROFILE`)
1. IAM Instance Profile (if running on EC2)

Add the following labels to the service in Rancher:
* `io.rancher.container.create_agent: true`
* `io.rancher.container.agent.role: environment`

These labels will cause Rancher to provision an API key for this service and
create the `CATTLE_URL`, `CATTLE_ACCESS_KEY`, and `CATTLE_SECRET_KEY`
environment variables.

## Auto creating registry in Rancher

This tool allows for automatically defining the ECR registry in Rancher by
setting the `AUTO_CREATE` environment variable to `true`.
When enabled, if the updater does not find an existing registry in Rancher
for the ECR URL, then it will automatically create the registry with the
proper credentials.
Subsequent executions of the update will simply update the credentials in Rancher
per normal operation.

## Configuring alternative ECR registries

By default the updater will acquire login tokens for the default registry
associated with the AWS account for the credentials used to access the AWS API.
This can be modified by providing the `AWS_ECR_REGISTRY_IDS` environment
variable to the container.
The variable should contain a comma (`,`) separated listed of account IDs to
acquire tokens for.
When specified, only the accounts provided will be looked up.
Each account will return an authorization token that will be used to update
and associated registry in Rancher.

## Running container outside of Rancher

If you are running this container outside of a Rancher managed environment, then
you must provide the following environment variables in additional to the ones
above.
* `CATTLE_URL`
* `CATTLE_ACCESS_KEY`
* `CATTLE_SECRET_KEY`

```bash
$ docker run -d -e AWS_REGION=us-east-1 -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY -e CATTLE_URL=http://rancher.mydomain.com -e CATTLE_ACCESS_KEY=$CATTLE_ACCESS_KEY -e CATTLE_SECRET_KEY=$CATTLE_SECRET_KEY objectpartners/rancher-ecr-credentials:latest
```

## Notes

The AWS credentials must correspond to an IAM user that has permissions to call
the ECR `GetToken` API.
The application then parses the resulting response to retrieve the ECR registry
URL, username, and password.
The returned registry URL, is used to discover the corresponding registry in
Rancher.

Rancher stores registries by environment.
If multiple environments exists, one instance of this container must be run per
environment.
Rancher credentials are tied to an environment, so specifying them will indicate
which environment to update in Rancher.

__NOTE__: This application runs on a 6 hour loop. It's possible there could be a
slight gap where the credentials expire before this program updates them.
