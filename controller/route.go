package controller

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/category-product/:categoryProductId", GetCategoryProductById)
	router.POST("/category-product", CreateCategoryProduct)
	router.PUT("/category-product/:categoryProductId", UpdateCategoryProductById)
	router.DELETE("/category-product/:categoryProductId", DeleteCategoryProductById)

	router.Run(":8080")
}
