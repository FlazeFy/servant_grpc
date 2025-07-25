package repository

import (
	"context"
	"servant/internal/model"
)

type TempSensorRepository interface {
	CreateTempSensor(ctx context.Context, reading *model.SensorReading) error
	FindAllTempSensor(ctx context.Context) ([]model.SensorReading, error)
}
