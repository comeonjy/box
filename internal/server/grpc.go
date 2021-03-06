package server

import (
	"time"

	"github.com/comeonjy/go-kit/grpc/reloadconfig"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/comeonjy/box/api/v1"
	"github.com/comeonjy/box/configs"
	"github.com/comeonjy/box/internal/service"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

func NewGrpcServer(srv *service.BoxService, conf configs.Interface, logger *xlog.Logger) *grpc.Server {
	server := grpc.NewServer(
		grpc.ConnectionTimeout(2*time.Second),
		grpc.ChainUnaryInterceptor(
			xmiddleware.GrpcLogger(xenv.GetEnv(xenv.TraceName), logger), xmiddleware.GrpcValidate, xmiddleware.GrpcRecover(logger), xmiddleware.GrpcAuth),
	)
	v1.RegisterBoxServer(server, srv)
	reloadconfig.RegisterReloadConfigServer(server, reloadconfig.NewServer(conf))
	return server
}
