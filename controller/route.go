package controller

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	// category-products
	router.GET("/category-products/:categoryProductId", GetCategoryProductById)
	router.POST("/category-products", CreateCategoryProduct)
	router.PUT("/category-products/:categoryProductId", UpdateCategoryProductById)
	router.DELETE("/category-products/:categoryProductId", DeleteCategoryProductById)
	router.GET("/category-products", GetAllCategoryProduct)
	router.GET("/category-products/:categoryProductId/products", GetAllProductByCategoryProductId)

	// provinces
	router.GET("/provinces/:provinceId", GetProvinceById)
	router.POST("/provinces", CreateProvince)
	router.PUT("/provinces/:provinceId", UpdateProvinceById)
	router.DELETE("/provinces/:provinceId", DeleteProvinceById)
	router.GET("/provinces",GetProvinces)

	router.Run(":8080")
}
