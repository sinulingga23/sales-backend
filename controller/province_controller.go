package controller

import (
	"fmt"
	"log"
	"time"
	"strings"
	"strconv"
	"net/http"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func GetProvinceById(c *gin.Context) {
	provinceId := 0

	provinceId, err := strconv.Atoi(c.Param("provinceId"));
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid",
			Errors:		"Bad Request",
		})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId);
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	fmt.Sprintf("%s", err),
			Errors:		"Not Found",
		})
		return
	}

	if isThere {
		currentProvince, err := provinceModel.FindProvinceById(provinceId);
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode: 	http.StatusInternalServerError,
				Message: 	"Somethings wrong!",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if currentProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, response.ResponseProvince {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the province",
				Province:	*currentProvince,
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

func CreateProvince(c *gin.Context) {
	requestProvince := model.Province{}

	err := c.Bind(&requestProvince)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message: "Invalid Request",
			Errors: "Bad Request",
		})
		return
	}

	if strings.Trim(requestProvince.Province, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message: 	"Province name can't be empty",
		})
		return
	} else {
		requestProvince.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		createdProvince, err := requestProvince.SaveProvince()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
				Errors:		fmt.Sprintf("%v", err),
			})
			return
		}

		if createdProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, response.ResponseProvince {
				StatusCode:	http.StatusOK,
				Message:	"Success to create province.",
				Province:	*createdProvince,
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

func UpdateProvinceById(c *gin.Context) {
	provinceId := 0
	requestProvince := model.Province{}

	err := c.Bind(&requestProvince);
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Somethings wrong!",
		})
		return
	}

	provinceId, err = strconv.Atoi(c.Param("provinceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	if provinceId != requestProvince.ProvinceId {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid format!",
		})
		return
	}

	if provinceId == 0 || requestProvince.ProvinceId == 0 {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"ProvinceId can't be zero",
		})
		return
	}

	if requestProvince.Province == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Province can't be empty"
		})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	fmt.Sprintf("%s", err),
			Errors:		"Not Found",
		})
		return
	}

	if isThere {
		currentProvince, err := provinceModel.FindProvinceById(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
				StatusCode:	http.StatusInternalServerError,
				Message:	fmt.Sprintf("%s", err),
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
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
			})
			return
		}

		if updatedProvince != (&model.Province{}) {
			c.JSON(http.StatusOK, response.ResponseProvince {
				StatusCode:	http.StatusOK
				Message:	"Success to update the province",
				Province:	*updatedProvince,})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}

func DeleteProvinceById(c *gin.Context) {
	provinceId := 0

	provinceId, err := strconv.Atoi(c.Param("provinceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request"})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	fmt.Sprintf("%s", err),
			Errors:		"Not Found",
		})
		return
	}

	if isThere {
		isDeleted, err := provinceModel.DeleteProvinceById(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
				Errors:		fmt.Sprintf("%s", err)
			})
			return
		}

		if isDeleted {
			c.JSON(http.StatusOK, response.ResponseGeneric {
				StatusCode:	http.StatusOK,
				Message:	"Success to delete province.",
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
