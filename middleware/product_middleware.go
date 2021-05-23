package middleware

import (
	"log"
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"
	"strings"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func ValidateProduct(c *gin.Context) {
	categoryProductModel := model.CategoryProduct{}
	requestProduct := model.Product{}

	buf, _ := ioutil.ReadAll(c.Request.Body)
	currentCheck := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = currentCheck

	err := c.Bind(&requestProduct)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	// For used by another handler
	requestProductAgain := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = requestProductAgain

	isThereInvalid := false
	listInvalid := make(map[string]string)

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
	if requestProduct.Name == "" {
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

	if isThereInvalid {
		c.JSON(http.StatusBadRequest, response.ResponseInvalids {
			StatusCode:	http.StatusBadRequest,
			Message:	"Make sure the fields is valid",
			Invalid:	listInvalid,
		})
		c.Abort()
		return
	}
	c.Next()
}
