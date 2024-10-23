# oauth2-token-vault

[![codecov](https://codecov.io/gh/microserv-io/oauth2-token-vault/graph/badge.svg?token=5TTII2E9NM)](https://codecov.io/gh/microserv-io/oauth2-token-vault)
[![golangci-lint](https://github.com/microserv-io/oauth2-token-vault/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/microserv-io/oauth2-token-vault/actions/workflows/golangci-lint.yml)
[![Run Tests](https://github.com/microserv-io/oauth2-token-vault/actions/workflows/tests.yml/badge.svg)](https://github.com/microserv-io/oauth2-token-vault/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/microserv-io/oauth2-token-vault)](https://goreportcard.com/report/github.com/microserv-io/oauth2-token-vault)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/microserv-io/oauth2-token-vault)]( )
> [!WARNING]
>
> This repository is a work in progress and is not yet ready for use. APIs are subject to change and will not be
> backwards compatible until release.

<p align="center">
  <img width="256" height="256" src="docs/logo.png" alt="logo">
</p>

Standalone service that handles storage of OAuth2 credentials for multiple providers. Using this service, you can
support OAuth2.0 resource servers in your application without needing to build your own token storage mechanisms.

This module wraps the [oauth2](https://pkg.go.dev/golang.org/x/oauth2) package from the Go standard library, and
therefore fully implements the OAuth2 protocol when communicating with Authorization Servers.

For deployment, please check out our public charts repository [here](https://github.com/microserv-io/public-charts).

### Deploy on Microserv.io

> [!TIP]
> Microserv.io is still being build and is not yet available for public use. Once Microserv.io is live, you can single
> click deploy this and other utility services to your private Mesh.

## Usage guide

This section describes deploying your own instance of the server, handling configuration parameters and interacting with
the service from your other services.

### Supported database servers

- PostgreSQL 14.0 or later (recommended)
- SQLite

### Deployment configuration

The service is configured through a `config.yaml` placed in root folder of the application.
If you use our Helm chart, you can override set the config using the `values.yaml` file.

The following is an example of the configuration file.

```yaml
providers:
  - name: google
    client_id: "google-client-id"
    client_secret: "google-client-secret"
    auth_url: "https://accounts.google.com/o/oauth2/auth"
    token_url: "https://accounts.google.com/o/oauth2/token"
    redirect_url: "http://localhost:8080/oauth/google/callback"
    scopes:
      - "https://www.googleapis.com/auth/analytics"
```

You can use environment variables to override the configuration file. This is recommended for secrets. The following is
an
example of how to override the configuration file with environment variables:

```bash
export PROVIDERS__0__SECRET_ID=google-secret-id
```

You can leave providers empty to not use any preconfigured providers.

#### Dynamic provider provisioning

If you want to dynamically add providers, for example, let your users add their own providers (for example: your
end-users own GitHub App) you can use the `AddProvider` method on the gRPC Client. You can also dynamically change or
remove providers that you created through the API.

### Interaction with the service

#### Authorization flow

Below flow explains the authorization flow. The Token Vault is only responsible for communication with the Authorization
Server, and will not in any way fetch data from the Resource Server.

![authorization flow](docs/authorization-flow.png)

Communication is done over gRPC. The repository contains a Go client that can be used to interact with the service. You
can install it by running the following command:

```bash
go get github.com/microserv-io/oauth2-token-vault@latest
```

### Retrieving resource servers with credentials

To retrieve access credentials, you can use the `GetCredentialsForProvider` method on the gRPC Client.

We also provide a direct integration with `golang.org/x/oauth2` package. See the following example:

```go
package main

import (
	"context"
	"github.com/microserv-io/oauth2-token-vault/pkg/oauth2/tokensource"
	"golang.org/x/oauth2"
	"net/url"
)

func main() {
	endpoint, _ := url.Parse("localhost:8080")
	// Create a new token source factory
	factory := tokensource.NewFactory(tokensource.WithEndpoint(endpoint))

	// Create a new client with the token source for the providerObj and resource owner	
	client := oauth2.NewClient(context.Background(), factory.CreateTokenSource(context.TODO(), "google", "some-user-id"))

	// Use the client to make requests
	_, _ = client.Get("https://www.googleapis.com/auth/analytics")
}

```

Please see the [examples](/examples) folder for a simple example of how to interact with the service.

## Contributing

### Protobufs

The service uses protobufs to define the API. The protobufs are located in the `proto` directory.

To generate the Go code from the protobufs, you can use the `buf` tool. To install `buf`, run the following command:

```bash
brew install buf
```

To generate the Go code from the protobufs, run the following command:

```bash
buf generate
```

### Database migrations with Atlas

We use Atlas to manage the database migrations. To run the migrations, you can use the following command:

```
brew install ariga/tap/atlas
```

To run the migrations, you can use the following command:

```bash
atlas schema apply
```

To create a new migration, you can use the following command:

```bash
atlas migrate diff
```

### Running the service

To run the service, you can use the following command:

```bash
go run cmd/grpc
```
