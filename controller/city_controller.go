package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sinulingga23/sales-backend/model"
	"github.com/sinulingga23/sales-backend/response"
	"github.com/sinulingga23/sales-backend/utility"
)

func GetCityById(c *gin.Context) {
	cityId := 0

	cityId, err := strconv.Atoi(c.Param("cityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		log.Printf("%s", err)
		return
	}

	cityModel := model.City{}
	isThere, err := cityModel.IsCityExistsById(cityId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The City is not exists",
		})
		return
	} else if isThere {
		currentCity, err := cityModel.FindCityById(cityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     fmt.Sprintf("%s", err),
			})
			log.Printf("%s", err)
			return
		}

		if currentCity != (&model.City{}) {
			c.JSON(http.StatusOK, response.ResponseCity{
				StatusCode: http.StatusOK,
				Message:    "Success to get the city",
				City:       *currentCity,
			})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func CreateCity(c *gin.Context) {
	requestCity := model.City{}
	provinceModel := model.Province{}

	err := c.Bind(&requestCity)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     "Bad Request",
		})
		log.Printf("%s", err)
		return
	}

	if requestCity.ProvinceId < 0 {
		c.JSON(http.StatusNotFound, response.ResponseErrors{
			StatusCode: http.StatusNotFound,
			Message:    "Somethings wrong!",
			Errors:     fmt.Sprintf("The Province with id %d is not exists.", requestCity.ProvinceId),
		})
		return
	}
	isThereProvince, err := provinceModel.IsProvinceExistsById(requestCity.ProvinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThereProvince {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The province is not exists.",
		})
		return
	} else if isThereProvince {
		if strings.Trim(requestCity.City, " ") == "" {
			c.JSON(http.StatusBadRequest, response.ResponseGeneric{
				StatusCode: http.StatusBadRequest,
				Message:    "City name can't be empty",
			})
			log.Printf("%s", err)
			return
		} else {
			requestCity.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
			createdCity, err := requestCity.SaveCity()
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseErrors{
					StatusCode: http.StatusInternalServerError,
					Message:    "Somethings wrong!",
					Errors:     fmt.Sprintf("%v", err),
				})
				log.Printf("%s", err)
				return
			}

			if createdCity != (&model.City{}) {
				c.JSON(http.StatusOK, response.ResponseCity{
					StatusCode: http.StatusOK,
					Message:    "Success to create a city.",
					City:       *createdCity,
				})
				log.Printf("%s", err)
				return
			}
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func UpdateCityById(c *gin.Context) {
	cityId := 0
	requestCity := model.City{}
	provinceModel := model.Province{}

	err := c.Bind(&requestCity)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Somethings wrong!",
		})
		log.Printf("%s", err)
		return
	}

	cityId, err = strconv.Atoi(c.Param("cityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Somethings wrong!",
		})
		log.Printf("%s", err)
		return
	}

	if cityId != requestCity.CityId || (cityId <= 0 || requestCity.CityId <= 0) {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid format!",
		})
		return
	}

	if strings.Trim(requestCity.City, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "City can't be empty",
		})
		return
	}

	isThereProvince, err := provinceModel.IsProvinceExistsById(requestCity.ProvinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThereProvince {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The province is not exists.",
		})
		return
	} else if isThereProvince {
		cityModel := model.City{}
		isThereCity, err := cityModel.IsCityExistsById(requestCity.CityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if !isThereCity {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "The City is not exists",
			})
			return
		} else if isThereCity {
			currentCity, err := cityModel.FindCityById(cityId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
					StatusCode: http.StatusInternalServerError,
					Message:    fmt.Sprintf("%s", err),
				})
				log.Printf("%s", err)
				return
			}

			if currentCity.Audit.CreatedAt != requestCity.Audit.CreatedAt {
				c.JSON(http.StatusBadRequest, response.ResponseGeneric{
					StatusCode: http.StatusBadRequest,
					Message:    "Invalid field createdAt",
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
				c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
					StatusCode: http.StatusInternalServerError,
					Message:    "Somethings wrong!",
				})
				log.Printf("%s", err)
				return
			}

			if updatedCity != (&model.City{}) {
				c.JSON(http.StatusOK, response.ResponseCity{
					StatusCode: http.StatusOK,
					Message:    "Success to update the city",
					City:       *updatedCity,
				})
				log.Printf("%s", err)
				return
			}
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func DeleteCityById(c *gin.Context) {
	cityId := 0

	cityId, err := strconv.Atoi(c.Param("cityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		log.Printf("%s", err)
		return
	}

	cityModel := model.City{}
	isThere, err := cityModel.IsCityExistsById(cityId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The City is not exists",
		})
		return
	} else if isThere {
		isDeleted, err := cityModel.DeleteCityById(cityId)
		if err != nil {
			log.Printf("%s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if !isDeleted {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "The City is not exists",
			})
		} else if isDeleted {
			c.JSON(http.StatusOK, response.ResponseGeneric{
				StatusCode: http.StatusOK,
				Message:    "Success to delete the city.",
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func GetCities(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	cityModel := model.City{}

	page := 0
	page, err := strconv.Atoi(requestPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	limit := 0
	limit, err = strconv.Atoi(requestLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	numberRecords, err := cityModel.GetNumberRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     "Internal Error",
		})
		return
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"cities"}, &page, &limit, numberRecords)
	offset := limit * (page - 1)

	cities, err := cityModel.FindAllCity(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if len(cities) != 0 {
		c.JSON(http.StatusOK, response.ResponseCities{
			StatusCode: http.StatusOK,
			Message:    "Success to get the cities",
			Cities:     cities,
			InfoPagination: response.InfoPagination{
				CurrentPage:  page,
				RowsEachPage: limit,
				TotalPages:   totalPages,
			},
			NextPage: nextPage,
			PrevPage: prevPage,
		})
		return
	}

	c.JSON(http.StatusNotFound, response.ResponseGeneric{
		StatusCode: http.StatusNotFound,
		Message:    "The cities is empty",
	})
	return
}

func GetSubDistrictsByCityId(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	cityId := 0

	page := 0
	page, err := strconv.Atoi(requestPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	limit := 0
	limit, err = strconv.Atoi(requestLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "The parameters invalid",
			Errors:     "Not Valid",
		})
		return
	}

	cityId, err = strconv.Atoi(c.Param("cityId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	cityModel := model.City{}
	isThere, err := cityModel.IsCityExistsById(cityId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The City is not exists",
		})
		return
	} else if isThere {
		subDistrictModel := model.SubDistrict{}
		numberRecordsSubDistrict, err := subDistrictModel.GetNumberRecordsByCityId(cityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     "Internal Error",
			})
			return
		}

		nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"provinces", strconv.Itoa(cityId), "cities"}, &page, &limit, numberRecordsSubDistrict)
		offset := limit * (page - 1)

		subDistricts, err := cityModel.FindAllSubDistrictByCityId(cityId, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if len(subDistricts) != 0 {
			c.JSON(http.StatusOK, response.ResponseSubDistrictsByCityId{
				StatusCode:   http.StatusOK,
				Message:      "Success to get the sub-districts by cityId",
				CityId:       cityId,
				SubDistricts: subDistricts,
				InfoPagination: response.InfoPagination{
					CurrentPage:  page,
					RowsEachPage: limit,
					TotalPages:   totalPages,
				},
				NextPage: nextPage,
				PrevPage: prevPage,
			})
			return
		} else {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "Can't found the sub-districts by cityId",
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}
