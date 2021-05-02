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

func UpdateCityById(c *gin.Context) {
	cityId := 0
	requestCity := model.City{}
	provinceModel := model.Province{}

	err := c.Bind(&requestCity)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Somethings wrong!",
		})
		log.Print("%s", err)
		return
	}

	cityId, err = strconv.Atoi(c.Param("cityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Somethings wrong!",
		})
		log.Printf("%s", err)
		return
	}

	if cityId != requestCity.CityId || (cityId <= 0 || requestCity.CityId <= 0){
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid format!",
		})
		return
	}

	if strings.Trim(requestCity.City, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"City can't be empty",
		})
		return
	}

	isThereProvince, err := provinceModel.IsProvinceExistsById(requestCity.ProvinceId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	fmt.Sprintf("%s", err),
			Errors:		"Not Found",
		})
		log.Printf("%s", err)
		return
	}

	if isThereProvince {
		cityModel := model.City{}
		isThereCity, err := cityModel.IsCityExistsById(requestCity.CityId)
		if err != nil {
			c.JSON(http.StatusNotFound, response.ResponseErrors {
				StatusCode:	http.StatusNotFound,
				Message:	fmt.Sprintf("%s", err),
				Errors:		"Not Found",
			})
			log.Printf("%s", err)
			return
		}

		if isThereCity {
			currentCity, err := cityModel.FindCityById(cityId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
					StatusCode:	http.StatusInternalServerError,
					Message:	fmt.Sprintf("%s", err),
				})
				log.Printf("%s", err)
				return
			}

			if currentCity.Audit.CreatedAt != requestCity.Audit.CreatedAt {
				c.JSON(http.StatusBadRequest, response.ResponseGeneric {
					StatusCode:	http.StatusBadRequest,
					Message:	"Invalid field createdAt",
				})
				return
			}

			currentCity.CityId = requestCity.CityId
			currentCity.ProvinceId = requestCity.ProvinceId
			currentCity.City = requestCity.City
			currentCity.Audit.CreatedAt = requestCity.Audit.CreatedAt
			timestamp := time.Now().Format("2006-01-02 15:05:03")
			currentCity.Audit.UpdatedAt = &timestamp


			updatedCity, err := currentCity.UpdateCityById(cityId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
					StatusCode:	http.StatusInternalServerError,
					Message:	"Somethings wrong!",
				})
				log.Printf("%s", err)
				return
			}

			if updatedCity != (&model.City{}) {
				c.JSON(http.StatusOK, response.ResponseCity {
					StatusCode:	http.StatusOK,
					Message:	"Success to update the city",
					City:		*updatedCity,
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

func DeleteCityById(c *gin.Context) {
	cityId := 0

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
		isDeleted, err := cityModel.DeleteCityById(cityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
				Errors:		fmt.Sprintf("%s", err),
			})
			log.Printf("%s", err)
			return
		}

		if isDeleted {
			c.JSON(http.StatusOK, response.ResponseGeneric {
				StatusCode:	http.StatusOK,
				Message:	"Success to delete the city.",
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
