package controller

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"

	"sales-backend/model"
	"github.com/gin-gonic/gin"
)


func GetCategoryProductById(c *gin.Context) {
	categoryProductId := c.Param("categoryProductId")

	if strings.Trim(categoryProductId," ") == "" {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusBadRequest, "CategoryProductId can't be empty"})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusNotFound, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
		}{http.StatusNotFound, fmt.Sprintf("%s", err)})
		return
	}

	if isThere {
		currentModel, err := categoryProductModel.FindCategoryProductById(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string	`json:"message"`
			}{http.StatusInternalServerError, fmt.Sprintf("%s", err)})
			return
		}

		if currentModel != (&model.CategoryProduct{}) {
			c.JSON(http.StatusOK, struct {
				StatusCode	int 			`json:"statusCode"`
				Message		string			`json:"message"`
				Data		model.CategoryProduct	`json:"categoryProduct"`
			}{http.StatusOK, "success to get data", *currentModel})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int 	`json:"statusCode"`
		Message		string	`json:"message"`
	}{http.StatusInternalServerError, "Somethings wrong!"})
	return
}


func CreateCategoryProduct(c *gin.Context) {
	requestCategoryProduct := model.CategoryProduct{}

	if err := c.Bind(&requestCategoryProduct); err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusInternalServerError, "Somethings wrong!"})
		return
	}

	if strings.Trim(requestCategoryProduct.Category, " ") == "" {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusBadRequest, "Category name can't be empty"})
		return
	} else {
		requestCategoryProduct.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		create, err := requestCategoryProduct.SaveCategoryProduct()
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
			}{http.StatusInternalServerError, fmt.Sprintf("%v", err)})
			return
		}

		if create != (&model.CategoryProduct{}) {
			c.JSON(http.StatusOK, struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
			}{http.StatusOK, "Success to create category product"})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int	`json:"statusCode"`
		Message		string	`json:"message"`
	}{http.StatusOK, "Somethings wrong!"})
	return
}

func UpdateCategoryProductById(c *gin.Context) {
	categoryProductId := c.Param("categoryProductId")
	requestCategoryProduct := model.CategoryProduct{}

	if err := c.Bind(&requestCategoryProduct); err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusInternalServerError, "Somethings wrong!"})
		return
	}

	if categoryProductId != requestCategoryProduct.CategoryProductId  {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusBadRequest, "Invalid format!"})
		return
	}

	if strings.Trim(categoryProductId, " ") == "" || strings.Trim(requestCategoryProduct.CategoryProductId, " ") == "" {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusBadRequest, "CategoryProductId can't be empty"})
		return
	} else if strings.Trim(requestCategoryProduct.Category, " ") == "" {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusBadRequest, "Category name can't be empty"})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusNotFound, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
		}{http.StatusNotFound, fmt.Sprintf("%s", err)})
		return
	}

	if isThere {
		currentModel, err := categoryProductModel.FindCategoryProductById(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string	`json:"message"`
			}{http.StatusInternalServerError, fmt.Sprintf("%s", err)})
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
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string	`json:"message"`
			}{http.StatusInternalServerError, "Somethings wrong!"})
			return
		}

		if updatedModel != (&model.CategoryProduct{}) {
			c.JSON(http.StatusOK, struct {
				StatusCode	int			`json:"statusCode"`
				Message		string			`json:"message"`
				Data		model.CategoryProduct	`json:"categoryProduct"`
			}{http.StatusOK, "success to update data", *updatedModel})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int 	`json:"statusCode"`
		Message		string	`json:"message"`
	}{http.StatusInternalServerError, "Somethings wrong!"})
	return
}

func DeleteCategoryProductById(c *gin.Context) {
	categoryProductId := c.Param("categoryProductId")

	if strings.Trim(categoryProductId, " ") == "" {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message		string	`json:"message"`
		}{http.StatusBadRequest, "CategoryProduct can't be empty"})
		return
	}

	categoryProductModel := model.CategoryProduct{}
	isThere, err := categoryProductModel.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		c.JSON(http.StatusNotFound, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
		}{http.StatusNotFound, fmt.Sprintf("%s", err)})
		return
	}

	if isThere {
		isSuccess, err := categoryProductModel.DeleteCategoryProductById(categoryProductId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string 	`json:"message"`
			}{http.StatusInternalServerError, fmt.Sprintf("%s", err)})
			return
		}

		if isSuccess {
			c.JSON(http.StatusOK, struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string 	`json:"message"`
			}{http.StatusOK, "Success to delete category product"})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int 	`json:"statusCode"`
		Message		string 	`json:"message"`
	}{http.StatusInternalServerError, "Somethings wrong!"})
}

func GetAllCategoryProduct(c *gin.Context) {
	categoryProductModel := model.CategoryProduct{}

	listCategoryProduct, err := categoryProductModel.FindAllCategoryProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
		}{http.StatusInternalServerError, "Somethings wrong!"})
		return
	}

	if len(listCategoryProduct) != 0 {
		c.JSON(http.StatusOK, struct {
			StatusCode	int				`json:"statusCode"`
			Message		string				`json:"message"`
			Data		[]*model.CategoryProduct	`json:"listCategoryProduct"`
		}{http.StatusOK, "Success to get all category product", listCategoryProduct})
		return
	} else {
		c.JSON(http.StatusNotFound, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
		}{http.StatusNotFound, "List category product is empty"})
		return
	}
}
