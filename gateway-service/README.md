# Gateway Service

### Description
Service for proxy HTTP REST API to gRPC services.

### Stack
- Go 1.23 (gin-gonic)

### How to start (dev)
```shell
cp .env.example .env
```

```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

### How to start (prod)
```shell
cp .env.example .env
```

```shell
docker-compose -f docker-compose.yml up -d
```

### gRPC build
```shell
protoc --go_out=. --go-grpc_out=. <.proto>
```

### Ports:
- Redis port: -
- DB port: -
- DLV port: 5865
- gRPC port: -
- HTTP port: 8000