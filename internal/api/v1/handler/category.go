package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tobygin.com/learn-gin/utils"
)

type CategoryHandler struct {
}

var validCategory = map[string]bool{
	"php":    true,
	"python": true,
	"golang": true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoriesV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if err := utils.ValidationInList("category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Run ok",
		"category": category,
	})
}
