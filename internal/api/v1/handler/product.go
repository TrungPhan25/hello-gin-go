package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Run ok",
	})
}

func (p *ProductHandler) GetProductV1(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Run ok",
		"product_id": productID,
	})
}

func (p *ProductHandler) CreateProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create product ok",
	})
}

func (p *ProductHandler) UpdateProductV1(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Update product ok",
		"product_id": productID,
	})
}

func (p *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Delete product ok",
		"product_id": productID,
	})
}
