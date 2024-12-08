# Gateway Service

### Description
Service for proxy HTTP REST API to gRPC services.

### Stack
- Go 1.23 (gin-gonic)

### How to start (dev)
```cmd
cp .env.example .env
```

```cmd
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

### How to start (prod)
```cmd
cp .env.example .env
```

```cmd
docker-compose -f docker-compose.yml up -d
```

### Ports:
- Redis port: -
- DB port: -
- DLV port: 5865
- gRPC port: -
- HTTP port: 8000