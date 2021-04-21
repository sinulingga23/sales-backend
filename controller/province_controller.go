package controller

import (
	"fmt"
	"time"
	"strings"
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

func CreateProvince(c *gin.Context) {
	requestProvince := model.Province{}

	err := c.Bind(&requestProvince)
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message   	string	`json:"message"`
			Errors 		string 	`json:"errors"`
		}{http.StatusBadRequest, "Invalid request", "Bad Request"})
		return
	}

	if strings.Trim(requestProvince.Province, " ") == "" {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int	`json:"statusCode"`
			Message   	string	`json:"message"`
		}{http.StatusBadRequest, "Province name can't be empty"})
		return
	} else {
		requestProvince.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		createdProvince, err := requestProvince.SaveProvince()
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode	int	`json:"statusCode"`
				Message		string	`json:"message"`
				Errors 		string 	`json:"errors"`
			}{http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%v", err)})
			return
		}

		if createdProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, struct {
				StatusCode	int		`json:"statusCode"`
				Message 	string		`json:"message"`
				Data 		model.Province	`json:"province"`
			}{http.StatusOK, "Success to create province.", *createdProvince})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int 	`json:"statusCode"`
		Message 	string 	`json:"message"`
	}{http.StatusInternalServerError, "Somethings wrong!"})
	return
}

func DeleteProvinceById(c *gin.Context) {
	provinceId := 0

	provinceId, err := strconv.Atoi(c.Param("provinceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, struct {
			StatusCode	int 	`json:"statusCode"`
			Message 	string 	`json:"message"`
			Errors		string 	`json:"errors"`
		}{http.StatusBadRequest, "Invalid Request", "Bad Request"})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusNotFound, struct {
			StatusCode	int 	`json:"statusCode"`
			Message		string 	`json:"message"`
			Errors		string	`json:"errors"`
		}{http.StatusNotFound, fmt.Sprintf("%s", err), "Not Found"})
		return
	}

	if isThere {
		isDeleted, err := provinceModel.DeleteProvinceById(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct {
				StatusCode 	int 	`json:"statusCode"`
				Message 	string 	`json:"message"`
				Errors 		string 	`json:"errors"`
			}{http.StatusInternalServerError, "Somethings wrong!", fmt.Sprintf("%s", err)})
			return
		}

		if isDeleted {
			c.JSON(http.StatusOK, struct {
				StatusCode	int 	`json:"statusCode"`
				Message		string 	`json:"message"`
			}{http.StatusOK, "Success to delete province."})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, struct {
		StatusCode	int 	`json:"statusCode"`
		Message 	string 	`json:"message"`
	}{http.StatusInternalServerError, "Somethings wrong!"})
	return
}
