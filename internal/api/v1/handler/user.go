package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"tobygin.com/learn-gin/utils"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	search := ctx.Query("search")

	if erro := utils.VaildationRequired("search", search); erro != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": erro.Error(),
		})
		return
	}

	if erro := utils.ValidationStringLength("search", search, 3, 50); erro != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": erro.Error(),
		})
		return
	}

	if erro := utils.ValidationRegex("search", search, searchRegex, "Search query can only contain letters, numbers, and spaces"); erro != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": erro.Error(),
		})
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Limit must be a positive integer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"search":  search,
	})
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, erro := utils.ValidationPositiveInt("ID", idStr)
	if erro != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"user_id": id,
	})
}

// UUID
func (u *UserHandler) GetUserByUUIDV1(ctx *gin.Context) {
	idStr := ctx.Param("uuid")

	uid, erro := utils.ValidationUUID("UUID", idStr)
	if erro != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"uuid":    uid,
	})
}

var (
	slugRegex         = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
	slugRegexCompiled = regexp.MustCompile(slugRegex)
)

// Slug
func (u *UserHandler) GetUserBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Slug cannot be empty",
		})
		return
	}

	if erro := utils.ValidationRegex("slug", slug, slugRegexCompiled, "Slug must be lowercase letters, numbers, and hyphens only"); erro != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"slug":    slug,
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
