package service

import (
	"context"
	"encoding/json"

	v1 "github.com/comeonjy/box/api/v1"
	"github.com/comeonjy/box/internal/data"
	"github.com/comeonjy/go-kit/pkg/xerror"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (svc *BoxService) FormGet(ctx context.Context, in *v1.FormGetReq) (*v1.FormGetResp, error) {
	form, err := svc.formRepo.GetFormByUUID(in.GetUuid())
	if err != nil {
		return nil, xerror.New(xerror.SQLErr, err.Error())
	}
	contentList, err := svc.formRepo.ListContentByFormUUID(in.GetUuid())
	if err != nil {
		return nil, xerror.New(xerror.SQLErr, err.Error())
	}

	res := v1.FormGetResp{
		Form: &v1.FormStruct{
			FormUuid:  form.UUID,
			FormTitle: form.FormTitle,
			SubTitle:  form.SubTitle,
		},
	}
	items := make([]*v1.FormItems, 0)
	for _, v := range contentList {
		options := make([]*v1.FormItemContentOptions, 0)
		if err := json.Unmarshal([]byte(v.Options), &options); err != nil {
			svc.logger.Error(ctx, err.Error())
			continue
		}
		extend := v1.FormItemContentExtend{}
		if err := json.Unmarshal([]byte(v.Extend), &extend); err != nil {
			svc.logger.Error(ctx, err.Error())
			continue
		}
		userAnswer := v1.UserAnswer{}
		if err := json.Unmarshal([]byte(v.UserAnswer), &userAnswer); err != nil {
			svc.logger.Error(ctx, err.Error())
			continue
		}
		items = append(items, &v1.FormItems{
			Content: &v1.FormItemContent{
				ContentTitle: v.ContentTitle,
				ContentType:  v.ContentType,
				ContentUuid:  v.UUID,
				Options:      options,
				Extend:       &extend,
				UserAnswer:   &userAnswer,
			},
		})
	}
	res.Form.Items = items
	return &res, nil
}

func (svc *BoxService) FormListQuestionType(ctx context.Context, in *emptypb.Empty) (*v1.FormListQuestionTypeResp, error) {
	return &v1.FormListQuestionTypeResp{
		List: []*v1.FormItems{
			{
				QuestionTypeName: "单选题",
				Content: &v1.FormItemContent{
					ContentTitle: "请选择一个选项",
					ContentType:  "radio",
					Extend:       &v1.FormItemContentExtend{},
					UserAnswer:   &v1.UserAnswer{},
					Options: []*v1.FormItemContentOptions{
						{
							OptionType:  "text",
							OptionValue: "1",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "选项1",
							},
						},
						{
							OptionType:  "text",
							OptionValue: "2",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "选项2",
							},
						},
					},
				},
			},
			{
				QuestionTypeName: "多选题",
				Content: &v1.FormItemContent{
					ContentTitle: "请选择选项",
					ContentType:  "checkbox",
					Extend:       &v1.FormItemContentExtend{},
					UserAnswer:   &v1.UserAnswer{},
					Options: []*v1.FormItemContentOptions{
						{
							OptionType:  "text",
							OptionValue: "1",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "选项1",
							},
						},
						{
							OptionType:  "text",
							OptionValue: "2",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "选项2",
							},
						},
					},
				},
			},
			{
				QuestionTypeName: "下拉题",
				Content: &v1.FormItemContent{
					ContentTitle: "请选择一个选项",
					ContentType:  "select",
					Extend:       &v1.FormItemContentExtend{},
					UserAnswer:   &v1.UserAnswer{},
					Options: []*v1.FormItemContentOptions{
						{
							OptionType:  "text",
							OptionValue: "1",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "选项1",
							},
						},
						{
							OptionType:  "text",
							OptionValue: "2",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "选项2",
							},
						},
					},
				},
			},
			{
				QuestionTypeName: "填空题",
				Content: &v1.FormItemContent{
					ContentTitle: "请填写以下内容",
					ContentType:  "fillblank",
					Extend:       &v1.FormItemContentExtend{},
					UserAnswer:   &v1.UserAnswer{},
					Options: []*v1.FormItemContentOptions{
						{
							OptionType: "text",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "填空1",
							},
						},
						{
							OptionType: "text",
							OptionContent: &v1.FormItemContentOptionsContent{
								Text: "填空2",
							},
						},
					},
				},
			},
		},
	}, nil
}

func (svc *BoxService) FormSave(ctx context.Context, in *v1.FormSaveReq) (*emptypb.Empty, error) {
	form := &data.FormModel{
		UUID:      in.GetForm().FormUuid,
		FormTitle: in.GetForm().FormTitle,
		SubTitle:  in.GetForm().SubTitle,
	}
	if len(in.GetForm().FormUuid) > 0 {
		if err := svc.formRepo.UpdateForm(form); err != nil {
			return nil, xerror.New(xerror.SQLErr, err.Error())
		}
	} else {
		form.UUID = uuid.NewString()
		if err := svc.formRepo.CreateForm(form); err != nil {
			return nil, xerror.New(xerror.SQLErr, err.Error())
		}
	}
	if err := svc.saveContent(ctx, form.UUID, in.GetForm().GetItems()); err != nil {
		return nil, xerror.New(xerror.SQLErr, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (svc *BoxService) saveContent(ctx context.Context, formUUID string, items []*v1.FormItems) error {
	for k, v := range items {
		userAnswer, err := json.Marshal(v.GetContent().GetUserAnswer())
		if err != nil {
			return err
		}
		extend, err := json.Marshal(v.GetContent().GetExtend())
		if err != nil {
			return err
		}
		options, err := json.Marshal(v.GetContent().GetOptions())
		if err != nil {
			return err
		}
		content := &data.FormContentModel{
			UUID:         v.GetContent().GetContentUuid(),
			FormUUID:     formUUID,
			ContentTitle: v.GetContent().ContentTitle,
			ContentType:  v.GetContent().ContentType,
			Extend:       string(extend),
			UserAnswer:   string(userAnswer),
			Options:      string(options),
			Sort:         uint64(k),
		}
		if len(v.GetContent().ContentUuid) > 0 {
			if err := svc.formRepo.UpdateFormContent(content); err != nil {
				return err
			}
		} else {
			content.UUID = uuid.NewString()
			if err := svc.formRepo.CreateFormContent(content); err != nil {
				return err
			}
		}
	}
	return nil
}

func (svc *BoxService) FormList(ctx context.Context, in *v1.FormListReq) (*v1.FormListResp, error) {
	form, err := svc.formRepo.ListForm()
	if err != nil {
		return nil, xerror.New(xerror.SQLErr, err.Error())
	}
	res := make([]*v1.FormListResp_FormList, 0)
	for _, v := range form {
		res = append(res, &v1.FormListResp_FormList{
			Uuid:      v.UUID,
			FormTitle: v.FormTitle,
			SubTitle:  v.SubTitle,
		})
	}
	return &v1.FormListResp{List: res}, nil
}

func (svc *BoxService) FormAnswerSave(ctx context.Context, in *v1.FormSaveReq) (*emptypb.Empty, error) {
	for _, v := range in.GetForm().GetItems() {
		arrValue, err := json.Marshal(v.GetContent().GetUserAnswer().GetArrValue())
		if err != nil {
			return nil, xerror.New(xerror.SQLErr, err.Error())
		}
		if err := svc.formRepo.CreateFormAnswer(&data.FormAnswerModel{
			UserUUID:        "",
			FormUUID:        in.GetForm().GetFormUuid(),
			FormContentUUID: v.GetContent().ContentUuid,
			Select:          v.GetContent().GetUserAnswer().GetSelect(),
			Other:           v.GetContent().GetUserAnswer().GetOther(),
			ArrValue:        string(arrValue),
		}); err != nil {
			return nil, xerror.New(xerror.SQLErr, err.Error())
		}
	}
	return &emptypb.Empty{}, nil
}
