package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func GetCityById(c *gin.Context) {
	cityId := 0;

	cityId, err := strconv.Atoi(c.Param("cityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid",
			Errors:		"Bad Request",
		})
		log.Printf("%s", err)
		return
	}


	cityModel := model.City{}
	isThere, err := cityModel.IsCityExistsById(cityId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	fmt.Sprintf("%s", err),
			Errors:		"Not Found",
		})
		log.Printf("%s", err)
		return
	}

	if isThere {
		currentCity, err := cityModel.FindCityById(cityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
				Errors:		fmt.Sprintf("%s", err),
			})
			log.Printf("%s", err)
			return
		}

		if currentCity != (&model.City{}) {
			c.JSON(http.StatusOK, response.ResponseCity {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the city",
				City:		*currentCity,
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
