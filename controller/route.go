package controller

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	router.GET("/category-products/:categoryProductId", GetCategoryProductById)
	router.POST("/category-products", CreateCategoryProduct)
	router.PUT("/category-products/:categoryProductId", UpdateCategoryProductById)
	router.DELETE("/category-products/:categoryProductId", DeleteCategoryProductById)
	router.GET("/category-products", GetAllCategoryProduct)

	router.Run(":8080")
}
