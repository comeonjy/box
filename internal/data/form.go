package data

import (
	"time"

	"gorm.io/gorm"
)

type FormModel struct {
	Id        uint64 `gorm:"primarykey"`
	UUID      string `gorm:"uniqueIndex;type:varchar(40)"`
	FormTitle string `gorm:"type:varchar(200)"`
	SubTitle  string `gorm:"type:varchar(200)"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type FormAnswerModel struct {
	Id              uint64 `gorm:"primarykey"`
	UserUUID        string
	FormUUID        string
	FormContentUUID string
	Select          string
	Other           string
	ArrValue        string `gorm:"type:json"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type FormContentModel struct {
	Id           uint64 `gorm:"primarykey"`
	UUID         string `gorm:"uniqueIndex;type:varchar(40)"`
	FormUUID     string `gorm:"type:varchar(40)"`
	ContentTitle string
	ContentType  string
	Extend       string `gorm:"type:json"`
	UserAnswer   string `gorm:"type:json"`
	Options      string `gorm:"type:json"`
	Sort         uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type formRepo struct {
	data *Data
}

func NewFormRepo(data *Data) FormRepo {
	return &formRepo{data: data}
}

func (r formRepo) GetFormByUUID(uuid string) (*FormModel, error) {
	res := FormModel{}
	err := r.data.box.Model(&FormModel{}).Where("uuid = ?", uuid).Find(&res).Error
	return &res, err
}

func (r formRepo) ListForm() ([]FormModel, error) {
	res := make([]FormModel, 0)
	err := r.data.box.Model(&FormModel{}).Find(&res).Error
	return res, err
}

func (r formRepo) ListContentByFormUUID(formUUID string) ([]FormContentModel, error) {
	res := make([]FormContentModel, 0)
	err := r.data.box.Model(&FormContentModel{}).Where("form_uuid = ?", formUUID).Find(&res).Error
	return res, err
}

func (r formRepo) CreateForm(form *FormModel) error {
	return r.data.box.Model(&FormModel{}).Create(form).Error
}

func (r formRepo) UpdateForm(form *FormModel) error {
	return r.data.box.Model(&FormModel{}).Where("uuid = ?", form.UUID).Updates(form).Error
}

func (r formRepo) CreateFormContent(content *FormContentModel) error {
	return r.data.box.Model(&FormContentModel{}).Create(content).Error
}

func (r formRepo) UpdateFormContent(content *FormContentModel) error {
	return r.data.box.Model(&FormContentModel{}).Where("uuid = ?", content.UUID).Updates(content).Error
}

func (r formRepo) CreateFormAnswer(answer *FormAnswerModel) error {
	return r.data.box.Model(&FormAnswerModel{}).Create(answer).Error
}
