package controller

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"strings"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	productId := c.Param("productId")

	if len(strings.Trim(productId, " ")) == 0 {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"ProductId can't be empty",
		})
		return
	}

	productModel := model.Product{}
	isThere, err := productModel.IsProductExistsById(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle the request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The product is not exists.",
		})
		return
	} else if isThere {
		currentProduct, err := productModel.FindProductById(productId)
		if err != nil {
			log.Printf("%s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if currentProduct != (&model.Product{}) {
			c.JSON(http.StatusOK, response.ResponseProduct {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the product",
				Product:	*currentProduct,
			})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}

func CreateProduct(c *gin.Context) {
	requestProduct := model.Product{}

	err := c.Bind(&requestProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	isThereInvalid := false
	listInvalid := map[string]string{}
	categoryProductModel := model.CategoryProduct{}

	// validate the CategoryProductId field
	if len(strings.Trim(" ",requestProduct.CategoryProductId)) == 0 {
		isThereInvalid = true
		listInvalid["err_category_product_id"] = "The CategoryProductId can't be empty"
	} else {
		isThereCategoryProduct, err := categoryProductModel.IsCategoryProductExistsById(requestProduct.CategoryProductId)
		if err != nil {
			log.Printf("%s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request.",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if !isThereCategoryProduct {
			isThereInvalid = true
			listInvalid["err_category_product_id"] = "The CategoryProductId is not exists."
		}
	}

	// validate the Name field
	if len(strings.Trim(" ", requestProduct.Name)) == 0 {
		isThereInvalid = true
		listInvalid["err_name"] = "The Name can't be empty"
	}

	// validtae the Unit field
	if len(strings.Trim(" ", requestProduct.Unit)) == 0 {
		isThereInvalid = true
		listInvalid["err_unit"] = "The Unit can't be empty"
	}

	// validate the Price field
	if requestProduct.Price < 0 {
		isThereInvalid = true
		listInvalid["err_price"] = "The Price can't be negative"
	}

	// validate the Stock field
	if requestProduct.Stock < 0 {
		isThereInvalid = true
		listInvalid["err_stock"] = "The Stock can't be negative"
	}

	// No errors
	if !isThereInvalid {
		requestProduct.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		createdProduct, err := requestProduct.SaveProduct()
		if err != nil {
			log.Printf("%s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if createdProduct != (&model.Product{}) {
			c.JSON(http.StatusOK, response.ResponseProduct {
				StatusCode:	http.StatusOK,
				Message:	"Success to create the product.",
				Product:	*createdProduct,
			})
			return
		}
	} else if isThereInvalid {
		c.JSON(http.StatusBadRequest, response.ResponseInvalids {
			StatusCode:	http.StatusBadRequest,
			Message:	"Make sure the fields is valid",
			Invalids:	listInvalid,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}
