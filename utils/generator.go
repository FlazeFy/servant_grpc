package utils

import "math/rand/v2"

func GenerateRandomTemp() float64 {
	return 20.0 + rand.Float64()*15.0
}

func GenerateRandomHumid() float64 {
	return 30.0 + rand.Float64()*70.0
}
