package controller

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/category-product/:categoryProductId", GetCategoryProductById)
	router.POST("/category-product", CreateCategoryProduct)

	router.Run(":8080")
}
