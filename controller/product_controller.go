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
	// categoryProductModel := model.CategoryProduct{}
	requestProduct := model.Product{}
	err := c.Bind(&requestProduct)
	if err != nil {
		log.Printf("ini error %s", err)
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

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

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}

func UpdateProductById(c *gin.Context) {
	productId := c.Param("productId")
	requestProduct := model.Product{}

	if err := c.Bind(&requestProduct); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Somethings wrong!",
		})
		return
	}

	if productId != requestProduct.ProductId {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid format",
		})
		return
	}

	if len(strings.Trim(productId, " ")) == 0 || len(strings.Trim(requestProduct.ProductId, " ")) == 0 {
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
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if currentProduct == (&model.Product{}) {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
			})
			return
		}

		log.Printf("requestProduct: %v\n", requestProduct)
		currentProduct.ProductId = requestProduct.ProductId
		currentProduct.Name = requestProduct.Name
		currentProduct.Unit = requestProduct.Unit
		currentProduct.Price = requestProduct.Price
		// The client is not permitted to direct update the stock field,
		// If the client won update the stock, should by the AddStock field
		if currentProduct.Stock != requestProduct.Stock  {
			c.JSON(http.StatusBadRequest, response.ResponseGeneric {
				StatusCode:	http.StatusBadRequest,
				Message:	"Invalid Request",
			})
			return
		}
		if (requestProduct.AddStock > 0) {
			currentProduct.Stock = currentProduct.Stock + requestProduct.AddStock
		}
		timestamp := time.Now().Format("2006-01-02 15:05:03")
		currentProduct.Audit.UpdatedAt = &timestamp

		log.Printf("currentProduct: %v\n", currentProduct)
		log.Printf("productId: %s", productId)
		updatedProduct, err := currentProduct.UpdateProductById(productId)
		if err != nil {
			log.Printf("198: %s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
			})
			return
		}

		if updatedProduct != (&model.Product{}) {
			c.JSON(http.StatusOK, response.ResponseProduct {
				StatusCode:	http.StatusOK,
				Message:	"Success to update the product.",
				Product:	*updatedProduct,
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
