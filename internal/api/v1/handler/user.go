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

type ListUsersRequest struct {
	Page    int    `json:"page" binding:"required,gte=1"`
	Limit   int    `json:"limit" binding:"required,gt=1,lte=100"`
	Sort    string `json:"sort" binding:"omitempty,oneof=asc desc"`
	Search  string `json:"search" binding:"omitempty,max=100"`
	Display bool   `json:"display" binding:"omitempty"`
}

// type ListUsersFormRequest struct {
// 	Page   int    `form:"page" binding:"required,gte=1"`
// 	Limit  int    `form:"limit" binding:"required,gte=1,lte=100"`
// 	Sort   string `form:"sort" binding:"omitempty,oneof=asc desc"`
// 	Search string `form:"search" binding:"omitempty,max=100"`
// }

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	var req ListUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"page":    req.Page,
		// "limit":   req.Limit,
		// "sort":    req.Sort,
		// "search":  req.Search,
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

type PostUsersRequest struct {
	Display *bool  `json:"display" binding:"omitempty"`
	Name    string `json:"name" binding:"required,min=2,max=50"`
	Address string `json:"address" binding:"required,min=2,max=200"`
}

func (u *UserHandler) CreateUserV1(ctx *gin.Context) {
	var req PostUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	if req.Display != nil {
		defaultDisplay := false
		req.Display = &defaultDisplay
	}

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
