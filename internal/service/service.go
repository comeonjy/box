package service

import (
	"context"

	"github.com/comeonjy/go-kit/pkg/xerror"
	"github.com/comeonjy/go-kit/pkg/xredis"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	account "github.com/comeonjy/account/api/v1"
	v1 "github.com/comeonjy/box/api/v1"
	"github.com/comeonjy/box/configs"
	"github.com/comeonjy/box/internal/data"
	"github.com/comeonjy/go-kit/pkg/xlog"
)

var ProviderSet = wire.NewSet(NewBoxService)

type BoxService struct {
	v1.UnimplementedBoxServer
	conf       configs.Interface
	logger     *xlog.Logger
	formRepo   data.FormRepo
	accountSvc account.AccountClient
	redis      *redis.Client
}

func NewBoxService(conf configs.Interface, logger *xlog.Logger, formRepo data.FormRepo) *BoxService {
	accountDial, err := grpc.Dial(conf.Get().AccountGrpc, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &BoxService{
		conf:       conf,
		formRepo:   formRepo,
		logger:     logger,
		accountSvc: account.NewAccountClient(accountDial),
		redis:      xredis.New(conf.Get().RedisConf),
	}
}

func (svc *BoxService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *BoxService) Ping(ctx context.Context, in *emptypb.Empty) (*v1.Result, error) {
	var arr []string
	if err := svc.redis.Keys(ctx, "*").ScanSlice(&arr); err != nil {
		return nil, xerror.New(xerror.RedisErr, err.Error())
	}
	return &v1.Result{
		Code:    200,
		Message: svc.conf.Get().Mode,
	}, nil
}
