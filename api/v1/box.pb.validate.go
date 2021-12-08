// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/box.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Result) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}

// Validate checks the field values on FormListQuestionTypeResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormListQuestionTypeResp) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FormListQuestionTypeRespValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FormListQuestionTypeRespValidationError is the validation error returned by
// FormListQuestionTypeResp.Validate if the designated constraints aren't met.
type FormListQuestionTypeRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormListQuestionTypeRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormListQuestionTypeRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormListQuestionTypeRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormListQuestionTypeRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormListQuestionTypeRespValidationError) ErrorName() string {
	return "FormListQuestionTypeRespValidationError"
}

// Error satisfies the builtin error interface
func (e FormListQuestionTypeRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormListQuestionTypeResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormListQuestionTypeRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormListQuestionTypeRespValidationError{}

// Validate checks the field values on FormGetReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FormGetReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uuid

	return nil
}

// FormGetReqValidationError is the validation error returned by
// FormGetReq.Validate if the designated constraints aren't met.
type FormGetReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormGetReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormGetReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormGetReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormGetReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormGetReqValidationError) ErrorName() string { return "FormGetReqValidationError" }

// Error satisfies the builtin error interface
func (e FormGetReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormGetReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormGetReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormGetReqValidationError{}

// Validate checks the field values on FormGetResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FormGetResp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetForm()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormGetRespValidationError{
				field:  "Form",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FormGetRespValidationError is the validation error returned by
// FormGetResp.Validate if the designated constraints aren't met.
type FormGetRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormGetRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormGetRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormGetRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormGetRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormGetRespValidationError) ErrorName() string { return "FormGetRespValidationError" }

// Error satisfies the builtin error interface
func (e FormGetRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormGetResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormGetRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormGetRespValidationError{}

// Validate checks the field values on FormSaveReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FormSaveReq) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetForm()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormSaveReqValidationError{
				field:  "Form",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FormSaveReqValidationError is the validation error returned by
// FormSaveReq.Validate if the designated constraints aren't met.
type FormSaveReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormSaveReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormSaveReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormSaveReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormSaveReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormSaveReqValidationError) ErrorName() string { return "FormSaveReqValidationError" }

// Error satisfies the builtin error interface
func (e FormSaveReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormSaveReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormSaveReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormSaveReqValidationError{}

// Validate checks the field values on FormListReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FormListReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FormListReqValidationError is the validation error returned by
// FormListReq.Validate if the designated constraints aren't met.
type FormListReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormListReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormListReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormListReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormListReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormListReqValidationError) ErrorName() string { return "FormListReqValidationError" }

// Error satisfies the builtin error interface
func (e FormListReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormListReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormListReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormListReqValidationError{}

// Validate checks the field values on FormListResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FormListResp) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FormListRespValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FormListRespValidationError is the validation error returned by
// FormListResp.Validate if the designated constraints aren't met.
type FormListRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormListRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormListRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormListRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormListRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormListRespValidationError) ErrorName() string { return "FormListRespValidationError" }

// Error satisfies the builtin error interface
func (e FormListRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormListResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormListRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormListRespValidationError{}

// Validate checks the field values on FormStruct with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FormStruct) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FormTitle

	// no validation rules for SubTitle

	// no validation rules for FormUuid

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FormStructValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FormStructValidationError is the validation error returned by
// FormStruct.Validate if the designated constraints aren't met.
type FormStructValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormStructValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormStructValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormStructValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormStructValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormStructValidationError) ErrorName() string { return "FormStructValidationError" }

// Error satisfies the builtin error interface
func (e FormStructValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormStruct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormStructValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormStructValidationError{}

// Validate checks the field values on FormItems with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FormItems) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormItemsValidationError{
				field:  "Content",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for QuestionTypeName

	return nil
}

// FormItemsValidationError is the validation error returned by
// FormItems.Validate if the designated constraints aren't met.
type FormItemsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormItemsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormItemsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormItemsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormItemsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormItemsValidationError) ErrorName() string { return "FormItemsValidationError" }

// Error satisfies the builtin error interface
func (e FormItemsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormItems.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormItemsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormItemsValidationError{}

// Validate checks the field values on FormItemContent with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FormItemContent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ContentTitle

	// no validation rules for ContentType

	// no validation rules for ContentUuid

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FormItemContentValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetExtend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormItemContentValidationError{
				field:  "Extend",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUserAnswer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormItemContentValidationError{
				field:  "UserAnswer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FormItemContentValidationError is the validation error returned by
// FormItemContent.Validate if the designated constraints aren't met.
type FormItemContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormItemContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormItemContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormItemContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormItemContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormItemContentValidationError) ErrorName() string { return "FormItemContentValidationError" }

// Error satisfies the builtin error interface
func (e FormItemContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormItemContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormItemContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormItemContentValidationError{}

// Validate checks the field values on FormItemContentOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormItemContentOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OptionType

	// no validation rules for OptionValue

	if v, ok := interface{}(m.GetOptionContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormItemContentOptionsValidationError{
				field:  "OptionContent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FormItemContentOptionsValidationError is the validation error returned by
// FormItemContentOptions.Validate if the designated constraints aren't met.
type FormItemContentOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormItemContentOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormItemContentOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormItemContentOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormItemContentOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormItemContentOptionsValidationError) ErrorName() string {
	return "FormItemContentOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e FormItemContentOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormItemContentOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormItemContentOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormItemContentOptionsValidationError{}

// Validate checks the field values on UserAnswer with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserAnswer) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Other

	// no validation rules for Select

	return nil
}

// UserAnswerValidationError is the validation error returned by
// UserAnswer.Validate if the designated constraints aren't met.
type UserAnswerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserAnswerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserAnswerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserAnswerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserAnswerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserAnswerValidationError) ErrorName() string { return "UserAnswerValidationError" }

// Error satisfies the builtin error interface
func (e UserAnswerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserAnswer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserAnswerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserAnswerValidationError{}

// Validate checks the field values on FormItemContentExtend with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormItemContentExtend) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Require

	return nil
}

// FormItemContentExtendValidationError is the validation error returned by
// FormItemContentExtend.Validate if the designated constraints aren't met.
type FormItemContentExtendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormItemContentExtendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormItemContentExtendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormItemContentExtendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormItemContentExtendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormItemContentExtendValidationError) ErrorName() string {
	return "FormItemContentExtendValidationError"
}

// Error satisfies the builtin error interface
func (e FormItemContentExtendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormItemContentExtend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormItemContentExtendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormItemContentExtendValidationError{}

// Validate checks the field values on FormItemContentOptionsContent with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormItemContentOptionsContent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Text

	// no validation rules for Explain

	// no validation rules for Img

	return nil
}

// FormItemContentOptionsContentValidationError is the validation error
// returned by FormItemContentOptionsContent.Validate if the designated
// constraints aren't met.
type FormItemContentOptionsContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormItemContentOptionsContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormItemContentOptionsContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormItemContentOptionsContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormItemContentOptionsContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormItemContentOptionsContentValidationError) ErrorName() string {
	return "FormItemContentOptionsContentValidationError"
}

// Error satisfies the builtin error interface
func (e FormItemContentOptionsContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormItemContentOptionsContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormItemContentOptionsContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormItemContentOptionsContentValidationError{}

// Validate checks the field values on FormListResp_FormList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormListResp_FormList) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uuid

	// no validation rules for FormTitle

	// no validation rules for SubTitle

	return nil
}

// FormListResp_FormListValidationError is the validation error returned by
// FormListResp_FormList.Validate if the designated constraints aren't met.
type FormListResp_FormListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormListResp_FormListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormListResp_FormListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormListResp_FormListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormListResp_FormListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormListResp_FormListValidationError) ErrorName() string {
	return "FormListResp_FormListValidationError"
}

// Error satisfies the builtin error interface
func (e FormListResp_FormListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormListResp_FormList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormListResp_FormListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormListResp_FormListValidationError{}