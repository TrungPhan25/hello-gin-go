package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	v1handler "tobygin.com/learn-gin/internal/api/v1/handler"
	"tobygin.com/learn-gin/middleware"
	"tobygin.com/learn-gin/utils"
)

func main() {
	r := gin.Default()

	erro := godotenv.Load()
	if erro != nil {
		log.Printf("No find env file")
	}

	if err := utils.RegisterValidation(); err != nil {
		panic("Failed to register custom validation: " + err.Error())
	}

	r.Use(middleware.APIKeyMiddleware())

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			userHandle := v1handler.NewUserHandler()

			user.GET("", userHandle.GetUsersV1)
			user.GET("/:user_id", userHandle.GetUserV1)
			user.GET("/uuid/:uuid", userHandle.GetUserByUUIDV1)
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

	r.Run(":8081")
}
