package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error) gin.H {
	if validationErro, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationErro {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn " + e.Param()
			}
		}

		return gin.H{
			"errors": errors,
		}
	}
	return gin.H{
		"message": "yêu cầu không hợp lệ " + err.Error(),
	}
}
