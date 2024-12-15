## CV GENERATOR API | WIP

### Description
This project is a simple CV generator API, written on Go in microservice architecture. It is a simple REST API with gRPC services.

### Links:
- [github](https://github.com/m1n64/cv-generator-api)
- [API docs](https://api.resumego.online/docs/)
- [Base url](https://api.resumego.online)

### Checklist:
- [x] user-service (Authorization, token validation, user info) (gRPC)
- [x] main cv-service (List of CV\'s, CRUD for CV\'s) (gRPC)
- [ ] cv information-service (biography, name, position, location, photo, education, work experience and skills) (gRPC) 
- [ ] file storage-service (minio) (gRPC, RabbitMQ(?))
- [ ] cv PDF generator-service (cv PDF generator) (gRPC, RabbitMQ)
- [ ] gateway-service (proxy for services from gRPC to REST) (REST API)
- [x] swagger docs service (or container in gateway-service) (REST API)

### Startup (makefile)
Init (for create network):
```shell
make network
```
```shell
make up
```

### Startup (docker-compose):
```shell
docker network create cv-generator-network
```
```shell
docker-compose -f shared-network/docker-compose.yml up -d
```
```shell
cd user-service
```
```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```
```shell
cd ../gateway-service
```
```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```
and all the services will be started. (yep, in future i will add k8s support, because I know that this is a "crutch" solution).
