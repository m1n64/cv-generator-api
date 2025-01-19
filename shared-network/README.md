# Shared Network

### Description
General Shared network and main containers

### Stack:
- nginx
- RabbitMQ
- MinIO

### How to start
```shell
cp .env.example .env
```
```shell
docker-compose up -d
```

### Ports:
- nginx: 80/443
- RabbitMQ: 15672
- MinIO: 9000
