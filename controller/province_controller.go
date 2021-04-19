package controller

import (
	"fmt"
	"strconv"
	"net/http"

	"sales-backend/model"
	"github.com/gin-gonic/gin"
)

func GetProvinceById(c *gin.Context) {
	provinceId := 0

	provinceId, err := strconv.Atoi(c.Param("provinceId"));
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int 	`json:"statusCode"`
			Message 	string 	`json:"message"`
			Errors		string 	`json:"errors"`
		}{http.StatusBadRequest, "Invalid Request", "Bad Request"})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId);
	if err != nil {
		c.JSON(http.StatusNotFound, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
			Errors		string	`json:"errors"`
		}{http.StatusNotFound, fmt.Sprintf("%s", err), "Not Found"})
		return
	}

	if isThere {
		currentProvince, err := provinceModel.FindProvinceById(provinceId);
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode 	int 	`json:"statusCode"`
				Message 	string 	`json:"message"`
				Errors 		string 	`json:"errors"`
			}{http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err)})
			return
		}

		if currentProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, struct {
				StatusCode	int		`json:"statusCode"`
				Message 	string		`json:"message"`
				Data 		model.Province	`json:"province"`
			}{http.StatusOK, "Success to get province.", *currentProvince})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int 	`json:"statusCode"`
		Message 	string 	`json:"message"`
	}{http.StatusInternalServerError, "Somethings wrong!"})
	return
}

