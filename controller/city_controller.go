package controller

import (
	"fmt"
	"log"
	"time"
	"strings"
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

func CreateCity(c *gin.Context) {
	requestCity := model.City{}
	provinceModel := model.Province{}

	err := c.Bind(&requestCity)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message: 	"Invalid Request",
			Errors: 	"Bad Request",
		})
		log.Printf("%s", err)
		return
	}

	if requestCity.ProvinceId < 0 {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	"Somethings wrong!",
			Errors:		fmt.Sprintf("The Province with id %d is not exists.", requestCity.ProvinceId),
		})
		return
	}
	isThereProvince, err := provinceModel.IsProvinceExistsById(requestCity.ProvinceId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	"Somethings wrong!",
			Errors:		fmt.Sprintf("%s", err),
		})
		log.Printf("%s", err)
		return
	}

	if isThereProvince {
		if strings.Trim(requestCity.City, " ") == "" {
			c.JSON(http.StatusBadRequest, response.ResponseGeneric {
				StatusCode:	http.StatusBadRequest,
				Message: 	"City name can't be empty",
			})
			log.Printf("%s", err)
			return
		} else  {
			requestCity.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
			createdCity, err := requestCity.SaveCity()
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"Somethings wrong!",
					Errors:		fmt.Sprintf("%v", err),
				})
				log.Printf("%s", err)
				return
			}

			if createdCity != (&model.City{}) {
				c.JSON(http.StatusOK, response.ResponseCity {
					StatusCode:	http.StatusOK,
					Message:	"Success to create a city.",
					City:		*createdCity,
				})
				log.Printf("%s", err)
				return
			}
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}
