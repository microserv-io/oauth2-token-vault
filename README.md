# oauth-credentials-server

> [!WARNING]
> 
> This repository is a work in progress and is not yet ready for use.

Standalone service that handles storage of OAuth2 credentials for multiple providers, allowing communication with other
services over gRPC.

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

You can use environment variables to override the configuration file. This is recommended for secrets. The following is an
example of how to override the configuration file with environment variables:

```bash
export PROVIDERS__0__SECRET_ID=google-secret-id
```

### Interaction with the service

Communication is done over gRPC. The repository contains a Go client that can be used to interact with the service. You
can install it by running the following command:

```bash
go get github.com/microserv-io/oauth-credentials-server@latest
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
