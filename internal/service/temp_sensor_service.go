package service

import (
	"context"
	"servant/servant/proto"
)

type TempSensorService interface {
	CreateTempSensor(ctx context.Context, req *proto.CreateTempSensorReq) (*proto.CreateTempSensorRes, error)
}
