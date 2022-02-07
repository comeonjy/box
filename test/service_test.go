package test

import (
	"context"
	"testing"

	v1 "github.com/comeonjy/box/api/v1"
	"github.com/comeonjy/go-kit/grpc/reloadconfig"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestReloadConfig(t *testing.T) {
	dial, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := reloadconfig.NewReloadConfigClient(dial)
	_, err = client.ReloadConfig(context.TODO(), &reloadconfig.Empty{})
	if err != nil {
		t.Error(err)
	}
}

func TestBoxService(t *testing.T) {
	dial, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := v1.NewBoxClient(dial)

	t.Run("ping", func(t *testing.T) {
		res, err := client.Ping(context.TODO(), &emptypb.Empty{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(res)
	})
}
