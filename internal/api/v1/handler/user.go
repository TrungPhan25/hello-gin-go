package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	search := ctx.Query("search")
	if search == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Search query cannot be empty",
		})
		return
	}

	if len(search) < 3 || len(search) > 50 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Search query must be between 3 and 50 characters long",
		})
		return
	}

	if !searchRegex.MatchString(search) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Search query can only contain letters, numbers, and spaces",
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
	})
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User ID must be greater than 0",
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

	_, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid UUID",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
		"uuid":    idStr,
	})
}

var slugRegex = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
var slugRegexCompiled = regexp.MustCompile(slugRegex)

// Slug
func (u *UserHandler) GetUserBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Slug cannot be empty",
		})
		return
	}

	if !slugRegexCompiled.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid slug format",
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
