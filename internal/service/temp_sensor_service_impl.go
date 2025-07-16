package service

import (
	"context"
	"time"

	"servant/internal/model"
	"servant/internal/repository"
	"servant/servant/proto"
)

type tempSensorService struct {
	repo repository.TempSensorRepository
}

func NewTempSensorService(repo repository.TempSensorRepository) TempSensorService {
	return &tempSensorService{repo: repo}
}

func (s *tempSensorService) CreateTempSensor(ctx context.Context, req *proto.CreateTempSensorReq) (*proto.CreateTempSensorRes, error) {
	// Model
	reading := &model.SensorReading{
		ID:          req.Id,
		Timestamp:   time.Now(),
		Temperature: req.Temperature,
		Humidity:    req.Humidity,
	}

	// Repository
	err := s.repo.CreateTempSensor(ctx, reading)
	if err != nil {
		return nil, err
	}

	return &proto.CreateTempSensorRes{Status: "success"}, nil
}

func (s *tempSensorService) GetAllTempSensors(ctx context.Context) (*proto.GetAllTempSensorRes, error) {
	// Repository
	sensors, err := s.repo.FindAllTempSensor(ctx)
	if err != nil {
		return nil, err
	}

	// Map to Proto
	var protoSensors []*proto.TempSensor
	for _, r := range sensors {
		protoSensors = append(protoSensors, &proto.TempSensor{
			Id:          r.ID,
			Timestamp:   r.Timestamp.Format(time.RFC3339),
			Temperature: r.Temperature,
			Humidity:    r.Humidity,
		})
	}

	return &proto.GetAllTempSensorRes{
		Status:  "success",
		Sensors: protoSensors,
	}, nil
}
