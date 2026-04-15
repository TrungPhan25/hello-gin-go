package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
	})
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"user_id": userID,
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
