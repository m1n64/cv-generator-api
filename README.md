## CV GENERATOR API | WIP

### Description
This project is a simple CV generator API, written on Go in microservice architecture. It is a simple REST API with gRPC services.

### Links:
- [github](https://github.com/m1n64/cv-generator-api)
- [API docs](https://api.resumego.online/docs/)
- [Base url](https://api.resumego.online)

### Checklist:
- [x] user-service (Authorization, token validation, user info) (gRPC)
- [ ] cv information-service (biography, name, position, location, photo, education, work experience and skills) (gRPC) 
- [ ] cv PDF generator-service (cv PDF generator) (gRPC)
- [ ] gateway-service (proxy for services from gRPC to REST) (REST API)
- [x] swagger docs service (or container in gateway-service) (REST API)

### Startup
```
docker network create cv-generator-network
```
```
docker-compose -f shared-network/docker-compose.yml up -d
```
```
cd user-service
```
```
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```
```
cd ../gateway-service
```
```
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```
and all the services will be started. (yep, in future i will add k8s support, because I know that this is a "crutch" solution).
