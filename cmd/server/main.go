package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"servant/config"
	"servant/internal/repository"
	"servant/internal/service"
	"servant/servant/proto"
)

type handler struct {
	proto.UnimplementedTempSensorServiceServer
	service service.TempSensorService
}

// Service
func (h *handler) CreateTempSensor(ctx context.Context, req *proto.CreateTempSensorReq) (*proto.CreateTempSensorRes, error) {
	return h.service.CreateTempSensor(ctx, req)
}
func (h *handler) GetAllTempSensor(ctx context.Context, req *proto.GetAllTempSensorReq) (*proto.GetAllTempSensorRes, error) {
	return h.service.GetAllTempSensors(ctx)
}

func main() {
	session := config.InitDatabase()
	defer session.Close()

	// Load Env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading ENV")
	}
	port := os.Getenv("PORT")

	// Dependency
	repoTempSensor := repository.NewTempSensorRepo(session)
	serviceTempSensor := service.NewTempSensorService(repoTempSensor)

	listener, _ := net.Listen("tcp", ":"+port)
	grpcServer := grpc.NewServer()
	proto.RegisterTempSensorServiceServer(grpcServer, &handler{service: serviceTempSensor})

	log.Println("gRPC server running at :" + port)
	grpcServer.Serve(listener)
}
