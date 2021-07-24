package controller

import (
	"sales-backend/middleware"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	// category-products
	router.GET("api/category-products/:categoryProductId", GetCategoryProductById)
	router.POST("api/category-products", middleware.ValidateTokenMiddleware(), CreateCategoryProduct)
	router.PUT("api/category-products/:categoryProductId", middleware.ValidateTokenMiddleware(), UpdateCategoryProductById)
	router.DELETE("api/category-products/:categoryProductId", middleware.ValidateTokenMiddleware(), DeleteCategoryProductById)
	router.GET("api/category-products", GetAllCategoryProduct)
	router.GET("api/category-products/:categoryProductId/products", GetAllProductByCategoryProductId)

	// provinces
	router.GET("api/provinces/:provinceId", GetProvinceById)
	router.POST("api/provinces", middleware.ValidateTokenMiddleware(), CreateProvince)
	router.PUT("api/provinces/:provinceId", middleware.ValidateTokenMiddleware(), UpdateProvinceById)
	router.DELETE("api/provinces/:provinceId", middleware.ValidateTokenMiddleware(), DeleteProvinceById)
	router.GET("api/provinces", GetProvinces)
	router.GET("api/provinces/:provinceId/cities", GetCitiesByProvinceId)

	// cities
	router.GET("api/cities/:cityId", GetCityById)
	router.POST("api/cities", middleware.ValidateTokenMiddleware(), CreateCity)
	router.PUT("api/cities/:cityId", middleware.ValidateTokenMiddleware(), UpdateCityById)
	router.DELETE("api/cities/:cityId", middleware.ValidateTokenMiddleware(), DeleteCityById)
	router.GET("api/cities", GetCities)
	router.GET("api/cities/:cityId/sub-districts", GetSubDistrictsByCityId)

	// sub-districts
	router.GET("api/sub-districts/:subDistrictId", GetSubDistrictById)
	router.POST("api/sub-districts", middleware.ValidateTokenMiddleware(), CreateSubDistrict)
	router.PUT("api/sub-districts/:subDistrictId", middleware.ValidateTokenMiddleware(), UpdateSubDistrictById)
	router.DELETE("api/sub-districts/:subDistrictId", middleware.ValidateTokenMiddleware(), DeleteSubDistrictById)
	router.GET("api/sub-districts", GetSubDistricts)

	// products
	router.GET("api/products/:productId", GetProductById)
	router.POST("api/products", middleware.ValidateTokenMiddleware(), middleware.ValidateProduct(), CreateProduct)
	router.PUT("api/products/:productId", middleware.ValidateTokenMiddleware(), middleware.ValidateProduct(), UpdateProductById)
	router.DELETE("api/products/:productId", middleware.ValidateTokenMiddleware(), DeleteProductById)
	router.GET("api/products", GetProducts)

	// users
	router.GET("api/users/:userId", GetUserById)
	router.POST("api/users", middleware.ValidateUser(), CreateUser)

	// login
	router.POST("api/login", BasicAuth)

	router.Run(":8080")
}
