package service

import (
	"context"

	"github.com/google/wire"
	"google.golang.org/grpc/metadata"

	v1 "github.com/comeonjy/box/api/v1"
	"github.com/comeonjy/box/configs"
	"github.com/comeonjy/box/internal/data"
	"github.com/comeonjy/go-kit/pkg/xlog"
)

var ProviderSet = wire.NewSet(NewSchedulerService)

type SchedulerService struct {
	v1.UnimplementedSchedulerServer
	conf     configs.Interface
	logger   *xlog.Logger
	workRepo data.WorkRepo
}

func NewSchedulerService(conf configs.Interface, logger *xlog.Logger, workRepo data.WorkRepo) *SchedulerService {
	return &SchedulerService{
		conf:     conf,
		workRepo: workRepo,
		logger:   logger,
	}
}

func (svc *SchedulerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *SchedulerService) Ping(ctx context.Context, in *v1.Empty) (*v1.Result, error) {
	return &v1.Result{
		Code:    200,
		Message: svc.conf.Get().Mode,
	}, nil
}