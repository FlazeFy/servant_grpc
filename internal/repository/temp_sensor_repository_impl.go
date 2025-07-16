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

func (r *tempSensorRepo) FindAllTempSensor(ctx context.Context) ([]model.SensorReading, error) {
	// Query
	query := `SELECT id, timestamp, temperature, humidity FROM sensor_readings`
	iter := r.db.Query(query).WithContext(ctx).Iter()

	// Model
	var readings []model.SensorReading
	var reading model.SensorReading

	// Scan
	for iter.Scan(&reading.ID, &reading.Timestamp, &reading.Temperature, &reading.Humidity) {
		readings = append(readings, reading)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return readings, nil
}
