package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tobygin.com/learn-gin/utils"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"limit":   limit,
	})
}

type GetUserByIDRequest struct {
	ID int `uri:"user_id" binding:"gt=0"`
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	var param GetUserByIDRequest
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"user_id": param.ID,
	})
}

func (u *UserHandler) CreateUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create user ok",
	})
}

func (u *UserHandler) UpdateUserV1(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user ok",
		"user_id": userID,
	})
}

func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete user ok",
		"user_id": userID,
	})
}

type GetuserByUUIDRequest struct {
	UUID string `uri:"uuid" binding:"uuid"`
}

func (u *UserHandler) GetUserByUUIDV1(ctx *gin.Context) {
	var param GetuserByUUIDRequest
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"uuid":    param.UUID,
	})
}
