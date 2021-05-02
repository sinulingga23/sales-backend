package controller

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	// category-products
	router.GET("api/category-products/:categoryProductId", GetCategoryProductById)
	router.POST("api/category-products", CreateCategoryProduct)
	router.PUT("api/category-products/:categoryProductId", UpdateCategoryProductById)
	router.DELETE("api/category-products/:categoryProductId", DeleteCategoryProductById)
	router.GET("api/category-products", GetAllCategoryProduct)
	router.GET("api/category-products/:categoryProductId/products", GetAllProductByCategoryProductId)

	// provinces
	router.GET("api/provinces/:provinceId", GetProvinceById)
	router.POST("api/provinces", CreateProvince)
	router.PUT("api/provinces/:provinceId", UpdateProvinceById)
	router.DELETE("api/provinces/:provinceId", DeleteProvinceById)
	router.GET("api/provinces",GetProvinces)
	router.GET("api/provinces/:provinceId/cities", GetCitiesByProvinceId)

	// cities
	router.GET("api/cities/:cityId", GetCityById)

	router.Run(":8080")
}
