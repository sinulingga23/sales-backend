package controller

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sinulingga23/sales-backend/middleware"
	"github.com/sinulingga23/sales-backend/pkg/monitoring"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()

	// cors
	router.Use(middleware.CORSMiddleware())

	// metric
	registry := prometheus.NewRegistry()
	monitoring.M = monitoring.NewMetric(registry)

	router.GET("/metrics", func(ctx *gin.Context) {
		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})

	// category-products
	router.GET("api/category-products/:categoryProductId", GetCategoryProductById)
	// router.POST("api/category-products", middleware.ValidateTokenMiddleware(), CreateCategoryProduct)
	router.POST("api/category-products", CreateCategoryProduct)
	router.PUT("api/category-products/:categoryProductId", middleware.ValidateTokenMiddleware(), UpdateCategoryProductById)
	router.DELETE("api/category-products/:categoryProductId", middleware.ValidateTokenMiddleware(), DeleteCategoryProductById)
	router.GET("api/category-products", GetAllCategoryProduct)
	router.GET("api/category-products/:categoryProductId/products", GetAllProductByCategoryProductId)

	// provinces
	router.GET("api/provinces/:provinceId", GetProvinceById)
	router.POST("api/provinces", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), CreateProvince)
	router.PUT("api/provinces/:provinceId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), UpdateProvinceById)
	router.DELETE("api/provinces/:provinceId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), DeleteProvinceById)
	router.GET("api/provinces", GetProvinces)
	router.GET("api/provinces/:provinceId/cities", GetCitiesByProvinceId)

	// cities
	router.GET("api/cities/:cityId", GetCityById)
	router.POST("api/cities", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), CreateCity)
	router.PUT("api/cities/:cityId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), UpdateCityById)
	router.DELETE("api/cities/:cityId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), DeleteCityById)
	router.GET("api/cities", GetCities)
	router.GET("api/cities/:cityId/sub-districts", GetSubDistrictsByCityId)

	// sub-districts
	router.GET("api/sub-districts/:subDistrictId", GetSubDistrictById)
	router.POST("api/sub-districts", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), CreateSubDistrict)
	router.PUT("api/sub-districts/:subDistrictId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), UpdateSubDistrictById)
	router.DELETE("api/sub-districts/:subDistrictId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), DeleteSubDistrictById)
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

	// roles
	router.GET("api/roles/:roleId", GetRoleById)
	router.POST("api/roles", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), CreateRole)
	router.PUT("api/roles/:roleId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), UpdateRole)
	router.DELETE("api/roles/:roleId", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), DeleteRole)
	router.GET("api/roles", GetRoles)

	// permissions
	router.GET("api/permissions/:permissionId", GetPermisisonById)
	router.POST("api/permissions", middleware.ValidateTokenMiddleware(), middleware.ValidateAdminMiddleware(), CreatePermission)

	// transaction
	router.GET("/api/transactions/:transactionId", GetTransactionById)
	router.POST("/api/transactions", middleware.ValidateTokenMiddleware(), CreateTransaction)

	router.Run(":8085")
}
