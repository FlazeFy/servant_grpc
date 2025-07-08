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
