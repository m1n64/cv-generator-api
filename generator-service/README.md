# CV PDF Generator Service

### Description
Service for generation PDF from CV all CV info.

### Stack:
- Go 1.23
- Postgres 15
- Redis

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
- Redis port: 6379
- DB port: 5432
- DLV port: 5868
- gRPC port: 50054
- HTTP port: -