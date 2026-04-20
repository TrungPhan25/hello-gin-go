package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) GetNewV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Run ok",
			"slug":    slug,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Run ok",
			"slug":    "No News",
		})
	}
}
