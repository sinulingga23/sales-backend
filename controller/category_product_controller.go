package controller

import (
	"fmt"
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
