package main

import (
	"context"
	"log"
	"os"

	"servant/servant/proto"
	"servant/utils"

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

	client := proto.NewTempSensorServiceClient(conn)

	// Service : Create Temp Sensor
	createResp, err := client.CreateTempSensor(context.Background(), &proto.CreateTempSensorReq{
		Id:          "TEST001",
		Temperature: float32(utils.GenerateRandomTemp()),
		Humidity:    float32(utils.GenerateRandomHumid()),
	})
	if err != nil {
		log.Fatalf("Error calling CreateTempSensor: %v", err)
	}
	log.Printf("CreateTempSensor response: %s", createResp.Status)

	// Service : Get All Temp Sensor
	getResp, err := client.GetAllTempSensor(context.Background(), &proto.GetAllTempSensorReq{})
	if err != nil {
		log.Fatalf("Error calling GetAllTempSensor: %v", err)
	}
	log.Printf("GetAllTempSensor response status: %s", getResp.Status)
	for _, sensor := range getResp.Sensors {
		log.Printf("Sensor - ID: %s, Timestamp: %s, Temp: %.2f, Humidity: %.2f",
			sensor.Id, sensor.Timestamp, sensor.Temperature, sensor.Humidity)
	}
}
