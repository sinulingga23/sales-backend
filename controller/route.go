package controller

import (
	"sales-backend/middleware"
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
	router.POST("api/cities", CreateCity)
	router.PUT("api/cities/:cityId", UpdateCityById)
	router.DELETE("api/cities/:cityId", DeleteCityById)
	router.GET("api/cities", GetCities)
	router.GET("api/cities/:cityId/sub-districts", GetSubDistrictsByCityId)

	// sub-districts
	router.GET("api/sub-districts/:subDistrictId", GetSubDistrictById)
	router.POST("api/sub-districts", CreateSubDistrict)
	router.PUT("api/sub-districts/:subDistrictId", UpdateSubDistrictById)
	router.DELETE("api/sub-districts/:subDistrictId", DeleteSubDistrictById)
	router.GET("api/sub-districts", GetSubDistricts)

	// products
	router.GET("api/products/:productId", GetProductById)
	router.POST("api/products", middleware.ValidateProduct(), CreateProduct)
	router.PUT("api/products/:productId", middleware.ValidateProduct(), UpdateProductById)
	router.DELETE("api/products/:productId", DeleteProductById)
	router.GET("api/products", GetProducts)

	// users
	router.GET("api/users/:userId", GetUserById)
	router.POST("api/users", middleware.ValidateUser(), CreateUser)

	router.Run(":8080")
}
