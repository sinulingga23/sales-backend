package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sales-backend/model"
	"sales-backend/response"
	"sales-backend/utility"

	"github.com/gin-gonic/gin"
)

func GetCategoryProductById(c *gin.Context) {
	categoryProductId := c.Param("categoryProductId")

	if strings.Trim(categoryProductId, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "CategoryProductId can't be empty",
		})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The category is not exists.",
		})
		return
	} else if isThere {
		currentModel, err := categoryProductModel.FindCategoryProductById(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("%s", err),
			})
			return
		}

		if currentModel != (&model.CategoryProduct{}) {
			c.JSON(http.StatusOK, response.ResponseCategoryProduct{
				StatusCode:      http.StatusOK,
				Message:         "success to get data",
				CategoryProduct: *currentModel,
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func CreateCategoryProduct(c *gin.Context) {
	requestCategoryProduct := model.CategoryProduct{}

	if err := c.Bind(&requestCategoryProduct); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request",
		})
		return
	}

	if strings.Trim(requestCategoryProduct.Category, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Category name can't be empty",
		})
		return
	} else {
		requestCategoryProduct.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		create, err := requestCategoryProduct.SaveCategoryProduct()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("%v", err),
			})
			return
		}

		if create != (&model.CategoryProduct{}) {
			c.JSON(http.StatusOK, response.ResponseGeneric{
				StatusCode: http.StatusOK,
				Message:    "Success to create category product",
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func UpdateCategoryProductById(c *gin.Context) {
	categoryProductId := c.Param("categoryProductId")
	requestCategoryProduct := model.CategoryProduct{}

	if err := c.Bind(&requestCategoryProduct); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Somethings wrong!",
		})
		return
	}

	if categoryProductId != requestCategoryProduct.CategoryProductId {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid format!",
		})
		return
	}

	if strings.Trim(categoryProductId, " ") == "" || strings.Trim(requestCategoryProduct.CategoryProductId, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "CategoryProductId can't be empty",
		})
		return
	} else if strings.Trim(requestCategoryProduct.Category, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Category name can't be empty"})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The category is not exists.",
		})
		return
	} else if isThere {
		currentModel, err := categoryProductModel.FindCategoryProductById(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("%s", err),
			})
			return
		}

		currentModel.CategoryProductId = requestCategoryProduct.CategoryProductId
		currentModel.Category = requestCategoryProduct.Category
		currentModel.Audit.CreatedAt = requestCategoryProduct.Audit.CreatedAt
		timestamp := time.Now().Format("2006-01-02 15:05:03")
		currentModel.Audit.UpdatedAt = &timestamp

		updatedModel, err := currentModel.UpdateCategoryProduct()
		if err != nil {
			log.Printf("%v", err)
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
			})
			return
		}

		if updatedModel != (&model.CategoryProduct{}) {
			c.JSON(http.StatusOK, response.ResponseCategoryProduct{
				StatusCode:      http.StatusOK,
				Message:         "success to update data",
				CategoryProduct: *updatedModel})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func DeleteCategoryProductById(c *gin.Context) {
	categoryProductId := c.Param("categoryProductId")

	if strings.Trim(categoryProductId, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "CategoryProduct can't be empty",
		})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The category is not exists.",
		})
		return
	} else if isThere {
		isSuccess, err := categoryProductModel.DeleteCategoryProductById(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
		}

		if !isSuccess {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "The category is not exists.",
			})
			return
		} else if isSuccess {
			c.JSON(http.StatusOK, response.ResponseGeneric{
				StatusCode: http.StatusOK,
				Message:    "Success to delete category product",
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func GetAllCategoryProduct(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	categoryProductModel := model.CategoryProduct{}

	page := 0
	limit := 0
	page, err := strconv.Atoi(requestPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	limit, err = strconv.Atoi(requestLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	numberRecords, err := categoryProductModel.GetNumberRecords()
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     "Internal Error",
		})
		return
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"category-products"}, &page, &limit, numberRecords)
	offset := limit * (page - 1)

	listCategoryProduct, err := categoryProductModel.FindAllCategoryProduct(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
		})
		return
	}

	if len(listCategoryProduct) != 0 {
		c.JSON(http.StatusOK, response.ResponseCategoryProducts{
			StatusCode:       http.StatusOK,
			Message:          "Success to get the category products",
			CategoryProducts: listCategoryProduct,
			InfoPagination: response.InfoPagination{
				CurrentPage:  page,
				RowsEachPage: limit,
				TotalPages:   totalPages,
			},
			NextPage: nextPage,
			PrevPage: prevPage,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "List category product is empty",
		})
		return
	}
}

func GetAllProductByCategoryProductId(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	categoryProductId := c.Param("categoryProductId")

	page := 0
	limit := 0
	page, err := strconv.Atoi(requestPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	limit, err = strconv.Atoi(requestLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The paramaters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	if strings.Trim(categoryProductId, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "CategoryProduct can't be empty",
		})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The category is not exists.",
		})
		return
	} else if isThere {
		productModel := model.Product{}
		numberRecordsProduct, err := productModel.GetNumberRecordsByCategoryProductId(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     "Internal Error",
			})
			return
		}

		log.Printf("req page: %d\n", page)
		log.Printf("req limit: %d\n", limit)
		nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"category-products", categoryProductId, "products"}, &page, &limit, numberRecordsProduct)
		offset := limit * (page - 1)

		listProduct, err := categoryProductModel.FindAllProductByCategoryProductId(categoryProductId, limit, offset)
		if err != nil {
			log.Printf("cpc: %v", err)
			c.JSON(http.StatusBadRequest, response.ResponseGeneric{
				StatusCode: http.StatusBadRequest,
				Message:    "Somethings wrong!",
			})
			return
		}

		if len(listProduct) != 0 {
			c.JSON(http.StatusOK, response.ResponseProductsByCategoryProductId{
				StatusCode:        http.StatusOK,
				Message:           "Success to get products by category id",
				CategoryProductId: categoryProductId,
				Products:          listProduct,
				InfoPagination: response.InfoPagination{
					CurrentPage:  page,
					RowsEachPage: limit,
					TotalPages:   totalPages,
				},
				NextPage: nextPage,
				PrevPage: prevPage,
			})
			return
		} else {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "Can't found the related list",
			})
			return
		}
	}
}
