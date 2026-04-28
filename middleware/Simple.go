package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before
		log.Printf("Before Check Middleware  ")
		// if 1 == 1 {
		// 	log.Printf("Abort Check Middleware  ")
		// 	c.Abort()
		// }
		c.Writer.Write([]byte("Before Simple Middleware\n"))
		c.Next()

		// After
		log.Printf("After Check Middleware  ")
		c.Writer.Write([]byte("After Simple Middleware\n"))
	}
}
