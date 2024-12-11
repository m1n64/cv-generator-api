# CV Information Service

### Description
Service for user's CV information (full name, position, biography, photo, skills, work employers, etc.).

### Stack:
- Go 1.23
- Postgres 15
- Redis

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
- Redis port: 6379
- DB port: 5432
- DLV port: 5867
- gRPC port: 50053
- HTTP port: -