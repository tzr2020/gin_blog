package validator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	uniTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/tzr2020/gin_blog/utils/errmsg"
)

// 验证结构体，将错误信息翻译为中文后返回
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := uniTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	// 注册翻译
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err: ", err)
		return "", errmsg.ERROR
	}

	// 设置验证结构体时，错误信息的字段名称
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})

	// 验证结构体，返回错误信息
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}

	return "", errmsg.SUCCESS
}
