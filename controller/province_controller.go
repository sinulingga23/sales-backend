package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sinulingga23/sales-backend/model"
	"github.com/sinulingga23/sales-backend/response"
	"github.com/sinulingga23/sales-backend/utility"

	"github.com/gin-gonic/gin"
)

func GetProvinceById(c *gin.Context) {

	provinceId := c.Param("provinceId")
	_, err := uuid.Parse(provinceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The province is not exists.",
		})
		return
	} else if isThere {
		currentProvince, err := provinceModel.FindProvinceById(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if currentProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, response.ResponseProvince{
				StatusCode: http.StatusOK,
				Message:    "Success to get the province",
				Province:   *currentProvince,
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

func CreateProvince(c *gin.Context) {
	requestProvince := model.Province{}

	err := c.Bind(&requestProvince)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     "Bad Request",
		})
		return
	}

	if strings.Trim(requestProvince.Province, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Province name can't be empty",
		})
		return
	} else {
		requestProvince.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		createdProvince, err := requestProvince.SaveProvince()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     fmt.Sprintf("%v", err),
			})
			return
		}

		if createdProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, response.ResponseProvince{
				StatusCode: http.StatusOK,
				Message:    "Success to create province.",
				Province:   *createdProvince,
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

func UpdateProvinceById(c *gin.Context) {
	requestProvince := model.Province{}

	err := c.Bind(&requestProvince)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Somethings wrong!",
		})
		return
	}

	provinceId := c.Param("provinceId")
	_, err = uuid.Parse(provinceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     "Bad Request",
		})
		return
	}

	if provinceId != requestProvince.ProvinceId {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid format!",
		})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The province is not exists.",
		})
		return
	} else if isThere {
		currentProvince, err := provinceModel.FindProvinceById(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
				StatusCode: http.StatusInternalServerError,
				Message:    fmt.Sprintf("%s", err),
			})
			return
		}

		currentProvince.ProvinceId = requestProvince.ProvinceId
		currentProvince.Province = requestProvince.Province
		currentProvince.Audit.CreatedAt = requestProvince.Audit.CreatedAt
		timestamp := time.Now().Format("2006-01-02 15:05:03")
		currentProvince.Audit.UpdatedAt = &timestamp

		updatedProvince, err := currentProvince.UpdateProvinceById(provinceId)
		if err != nil {
			log.Printf("%v", err)
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
			})
			return
		}

		if updatedProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, response.ResponseProvince{
				StatusCode: http.StatusOK,
				Message:    "Success to update the province",
				Province:   *updatedProvince,
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

func DeleteProvinceById(c *gin.Context) {

	provinceId := c.Param("provinceId")
	_, err := uuid.Parse(provinceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     "Bad Request"})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The province is not exists.",
		})
		return
	} else if isThere {
		isDeleted, err := provinceModel.DeleteProvinceById(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if !isDeleted {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "The province is not exists.",
			})
			return
		} else if isDeleted {
			c.JSON(http.StatusOK, response.ResponseGeneric{
				StatusCode: http.StatusOK,
				Message:    "Success to delete province.",
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

func GetProvinces(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	provinceModel := model.Province{}

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

	numberRecords, err := provinceModel.GetNumberRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     "Internal Error",
		})
		return
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"provinces"}, &page, &limit, numberRecords)
	offset := limit * (page - 1)

	provinces, err := provinceModel.FindAllProvince(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if len(provinces) != 0 {
		c.JSON(http.StatusOK, response.ResponseProvinces{
			StatusCode: http.StatusOK,
			Message:    "Success to get the provinces",
			Provinces:  provinces,
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
			Message:    "The provinces is empty",
		})
		return
	}
}

func GetCitiesByProvinceId(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")

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

	provinceId := c.Param("provinceId")
	_, err = uuid.Parse(provinceId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The province is not exists.",
		})
		return
	} else if isThere {
		cityModel := model.City{}
		numberRecordsCity, err := cityModel.GetNumberRecordsByProvinceId(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     "Internal Error",
			})
			return
		}

		nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"provinces", provinceId, "cities"}, &page, &limit, numberRecordsCity)
		offset := limit * (page - 1)

		cities, err := provinceModel.FindAllCityByProvinceId(provinceId, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if len(cities) != 0 {
			c.JSON(http.StatusOK, response.ResponseCitiesByProvinceId{
				StatusCode: http.StatusOK,
				Message:    "Success to get the cities by provinceId",
				ProvinceId: provinceId,
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
		} else {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "Can't found the cities by provinceId",
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
