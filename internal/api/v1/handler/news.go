package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tobygin.com/learn-gin/utils"
)

type NewsHandler struct {
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

type GetNewsBySlugRequest struct {
	Slug string `uri:"slug" binding:"slug,min=3,max=100"`
}

func (n *NewsHandler) GetNewV1(ctx *gin.Context) {
	var req GetNewsBySlugRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"m essage": "Run ok",
		"slug":     req.Slug,
	})
}
