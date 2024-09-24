# oauth-credentials-server

> [!WARNING]
> 
> This repository is a work in progress and is not yet ready for use.

Standalone service that handles storage of OAuth2 credentials for multiple providers, allowing communication with other
services over gRPC.

For deployment, please checkout our public charts repository [here](https://github.com/microserv-io/public-charts).

### Deploy on Microserv.io
> [!TIP]
> Microserv.io is still being build and is not yet available for public use. Once Microserv.io is live, you can single click deploy this and other utility services to your private Mesh.
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
