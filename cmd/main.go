package main

import (
	"ishlab_chiqarish/internal/config"
	"ishlab_chiqarish/internal/pkg/logger"
	"ishlab_chiqarish/internal/service"
	"ishlab_chiqarish/internal/storage/postgres"
	"log"
	"net"
	pb "ishlab_chiqarish/genproto/contract"
	pbs "ishlab_chiqarish/genproto/services"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().PRODUCTION_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	
	contractservice := service.NewProductionService(db, logger.NewLogger())
	service_employee := service.NewEmployeeServiceRepo(db, logger.NewLogger())

	server := grpc.NewServer()
	pb.RegisterContractServiceServer(server, contractservice)
	pbs.RegisterServicesServiceServer(server, service_employee)

	log.Printf("Server is listening on port %s\n", config.Load().PRODUCTION_SERVICE)
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}