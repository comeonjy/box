package consts

import (
	"github.com/comeonjy/go-kit/pkg/xenv"
)

var EnvMap = map[string]string{
	xenv.AppName:     "box",
	xenv.AppVersion:  "v1.0",
	xenv.ApolloAppID: "box",
	xenv.ApolloUrl:   "http://apollo.dev.jiangyang.me",
	"my_const":       "my_const_value",
}
