package model

import "time"

type SensorReading struct {
	ID          string
	Timestamp   time.Time
	Temperature float32
	Humidity    float32
}
