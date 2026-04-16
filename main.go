package main

import (
	v1handler "tobygin.com/learn-gin/internal/api/v1/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			userHandle := v1handler.NewUserHandler()

			user.GET("", userHandle.GetUsersV1)
			user.GET("/:id", userHandle.GetUserV1)
			user.GET("/admin/:uuid", userHandle.GetUserByUUIDV1)
			// user.GET("/admin/:slug", userHandle.GetUserBySlugV1)
			user.POST("", userHandle.CreateUserV1)
			user.PUT("/:id", userHandle.UpdateUserV1)
			user.DELETE("/:id", userHandle.DeleteUserV1)
		}

		category := v1.Group("/category")
		{
			categoryHandle := v1handler.NewCategoryHandler()

			category.GET(":category", categoryHandle.GetCategoryByCategoriesV1)
		}
	}

	r.Run(":8080")
}
