package repository

import (
	"context"
	"servant/internal/model"

	"github.com/gocql/gocql"
)

type tempSensorRepo struct {
	db *gocql.Session
}

func NewTempSensorRepo(session *gocql.Session) TempSensorRepository {
	return &tempSensorRepo{db: session}
}

func (r *tempSensorRepo) CreateTempSensor(ctx context.Context, reading *model.SensorReading) error {
	return r.db.Query(`
		INSERT INTO sensor_readings (id, timestamp, temperature, humidity)
		VALUES (?, ?, ?, ?)`,
		reading.ID, reading.Timestamp, reading.Temperature, reading.Humidity,
	).WithContext(ctx).Exec()
}
