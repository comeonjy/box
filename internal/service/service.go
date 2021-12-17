package service

import (
	"context"

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
	}
}

func (svc *BoxService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	if mdIn, ok := metadata.FromIncomingContext(ctx); ok {
		mdIn.Get("")
	}
	return ctx, nil
}

func (svc *BoxService) Ping(ctx context.Context, in *emptypb.Empty) (*v1.Result, error) {
	return &v1.Result{
		Code:    200,
		Message: svc.conf.Get().Mode,
	}, nil
}

type FormStruct struct {
	FormTitle string `json:"form_title"`
	SubTitle  string `json:"sub_title"`
	Items     []struct {
		Content struct {
			ContentTitle string `json:"content_title"`
			ContentType  string `json:"content_type"`
			Options      []struct {
				OptionType    string `json:"option_type"`
				OptionContent struct {
					Text    string `json:"text"`
					Explain string `json:"explain"`
					Img     string `json:"img"`
				} `json:"option_content"`
				OptionValue string `json:"option_value"`
			} `json:"options"`
			Extend struct {
				Require bool `json:"require"`
			} `json:"extend"`
			UserAnswer struct {
				Other    string   `json:"other"`
				Select   string   `json:"select"`
				ArrValue []string `json:"arr_value"`
			} `json:"user_answer"`
		} `json:"content"`
	} `json:"items"`
}
