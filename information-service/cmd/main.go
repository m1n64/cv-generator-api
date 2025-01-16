package main

import (
	"fmt"
	"google.golang.org/grpc"
	certificates "information-service/internal/certificates/grpc"
	handlers4 "information-service/internal/certificates/handlers"
	contacts "information-service/internal/contacts/grpc"
	handlers5 "information-service/internal/contacts/handlers"
	educations "information-service/internal/educations/grpc"
	handlers6 "information-service/internal/educations/handlers"
	experiences "information-service/internal/experiences/grpc"
	handlers7 "information-service/internal/experiences/handlers"
	health "information-service/internal/health/grpc"
	handlers8 "information-service/internal/health/handlers"
	information "information-service/internal/information/grpc"
	"information-service/internal/information/handlers"
	languages "information-service/internal/languages/grpc"
	handlers2 "information-service/internal/languages/handlers"
	skills "information-service/internal/skills/grpc"
	handlers3 "information-service/internal/skills/handlers"
	"information-service/pkg/containers"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("CV Information service started!")

	dependencies, _ := containers.InitializeDependencies()
	go containers.InitializeQueuesConsumer(dependencies)

	port := os.Getenv("SERVICE_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	cvInfoServer := handlers.NewCVInformationServiceServer(dependencies.CVInformationService, dependencies.FileService, dependencies.Logger)
	information.RegisterInformationServiceServer(grpcServer, cvInfoServer)

	langServer := handlers2.NewLanguageServiceServer(dependencies.LanguageService, dependencies.Logger)
	languages.RegisterLanguagesServiceServer(grpcServer, langServer)

	skillsServer := handlers3.NewSkillServiceServer(dependencies.SkillService, dependencies.Logger)
	skills.RegisterSkillsServiceServer(grpcServer, skillsServer)

	certsServer := handlers4.NewCertificateServiceServer(dependencies.CertificateService, dependencies.Logger)
	certificates.RegisterCertificatesServiceServer(grpcServer, certsServer)

	contactServer := handlers5.NewContactServiceServer(dependencies.ContactService, dependencies.Logger)
	contacts.RegisterContactsServiceServer(grpcServer, contactServer)

	educationServer := handlers6.NewEducationServiceServer(dependencies.EducationService, dependencies.Logger)
	educations.RegisterEducationServiceServer(grpcServer, educationServer)

	experienceServer := handlers7.NewWorkExperienceServiceServer(dependencies.WorkExperienceService, dependencies.Logger)
	experiences.RegisterExperiencesServiceServer(grpcServer, experienceServer)

	healthServer := handlers8.NewHealthServiceServer(dependencies.DB, dependencies.RedisClient)
	health.RegisterHealthServiceServer(grpcServer, healthServer)

	log.Printf("gRPC server is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
