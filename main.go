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
			user.GET("/:user_id", userHandle.GetUserV1)
			user.POST("", userHandle.CreateUserV1)
			user.PUT("/:user_id", userHandle.UpdateUserV1)
			user.DELETE("/:user_id", userHandle.DeleteUserV1)
		}

		news := v1.Group("/news")
		{
			newsHandler := v1handler.NewNewsHandler()

			news.GET("", newsHandler.GetNewV1)
			news.GET("/:slug", newsHandler.GetNewV1)
		}
	}

	r.Run(":8080")
}
