package data

import (
	"github.com/comeonjy/box/configs"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewWorkRepo)

type Data struct {
}

func NewData(cfg configs.Interface) *Data {
	return &Data{}
}