package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/demo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": []string{"Alice", "Bob", "Charlie"},
		})
	})

	r.GET("/user/:user_id", func(c *gin.Context) {
		user_id := c.Param("user_id")
		name_query := c.Query("name")
		c.JSON(200, gin.H{
			"data":    "Thoong tin user",
			"user_id": user_id,
			"name":    name_query,
		})
	})

	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"products": "Danh sach san pham",
		})
	})

	r.Run(":8080")

}
