package main

import (
	"context"
	"log"
	"os"

	"servant/servant/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading ENV")
	}
	port := os.Getenv("PORT")

	// Connect to Server
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Test Add Temp Sensor
	client := proto.NewTempSensorServiceClient(conn)
	resp, err := client.CreateTempSensor(context.Background(), &proto.CreateTempSensorReq{
		Id:          "TEST001",
		Temperature: 25.4,
		Humidity:    70.2,
	})
	if err != nil {
		log.Fatalf("Error calling CreateTempSensor: %v", err)
	}

	log.Printf("Response: %s", resp.Status)
}
