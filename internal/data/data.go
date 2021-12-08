package data

import (
	"github.com/comeonjy/box/configs"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmysql"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewFormRepo)

type Data struct {
	box *gorm.DB
}

func NewData(cfg configs.Interface, xlogger *xlog.Logger) *Data {
	db := xmysql.New(cfg.Get().MysqlConf, xlogger)
	if err := db.AutoMigrate(
		&FormModel{},
		&FormContentModel{},
		&FormAnswerModel{},
	); err != nil {
		panic(err)
	}
	return &Data{
		box: db,
	}
}

type FormRepo interface {
	GetFormByUUID(uuid string) (*FormModel, error)
	ListForm() ([]FormModel, error)
	ListContentByFormUUID(formUUID string) ([]FormContentModel, error)
	CreateForm(form *FormModel) error
	UpdateForm(form *FormModel) error
	CreateFormContent(content *FormContentModel) error
	UpdateFormContent(content *FormContentModel) error
	CreateFormAnswer(answer *FormAnswerModel) error
}
