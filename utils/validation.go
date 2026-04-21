package utils

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) gin.H {
	if validationErro, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationErro {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn " + e.Param()
			case "uuid":
				errors[e.Field()] = e.Field() + " phải là một UUID hợp lệ"
			case "slug":
				errors[e.Field()] = e.Field() + " phải là một slug hợp lệ chỉ chứa chữ thường, số và dấu gạch nối"
			// case "min":
			// 	errors[e.Field()] = e.Field() + " phải có độ dài tối thiểu là " + e.Param()
			case "max":
				errors[e.Field()] = e.Field() + " phải có độ dài tối đa là " + e.Param()
			}
		}

		return gin.H{
			"errors":  "Yêu cầu không hợp lệ",
			"details": validationErro.Error(),
		}
	}
	return gin.H{
		"message": "yêu cầu không hợp lệ",
	}
}

func RegisterValidation() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return errors.New("không thể đăng ký trình xác thực tùy chỉnh")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	return nil
}
