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
			Message:	"Province can't be empty",
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
				StatusCode:	http.StatusOK,
				Message:	"Success to update the province",
				Province:	*updatedProvince,
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
				Errors:		fmt.Sprintf("%s", err),
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

func GetProvinces(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	provinceModel := model.Province{}

	page := 0
	page, err := strconv.Atoi(requestPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"The parameters invalid",
			Errors:		"Not Valid",
		})
		return
	}

	limit := 0
	limit, err = strconv.Atoi(requestLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"The parameters invalid",
			Errors:		"Not Valid",
		})
		return
	}

	if page < 0 {
		page = 1
	}


	if limit < 0 {
		limit = 10
	} else if (limit > 25) {
		limit = 25
	}

	numberRecords, err := provinceModel.GetNumberRecords()

	totalPages := 0
	if totalPages = numberRecords / limit; numberRecords % limit != 0 {
		totalPages += 1
	}

	nextPage := fmt.Sprintf("api/provinces?page=%d&limit=%d", page+1, limit)
	prevPage := fmt.Sprintf("api/provinces?page=%d&limit=%d", page-1, limit)

	if (page+1) > totalPages {
		nextPage = ""
	} else if (page-1) < 1 {
		prevPage = ""
	}

	if (page >= 1 && limit >= numberRecords) {
		page = 1
		limit = numberRecords
		prevPage = ""
	}
	offset := limit * (page-1)

	provinces, err := provinceModel.FindAllProvince(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message: 	"Somethings wrong!",
			Errors: 	fmt.Sprintf("%s", err),
		})
		return
	}

	if len(provinces) != 0 {
		c.JSON(http.StatusOK, response.ResponseProvinces {
			StatusCode:	http.StatusOK,
			Message:	"Success to get the provinces",
			Provinces:	provinces,
			InfoPagination:	response.InfoPagination {
				CurrentPage:	page,
				RowsEachPage:	limit,
				TotalPages:	totalPages,
			},
			NextPage:	nextPage,
			PrevPage:	prevPage,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The provinces is empty",
		})
		return
	}
}

func GetCitiesByProvinceId(c *gin.Context) {
	provinceId := 0

	provinceId, err := strconv.Atoi(c.Param("provinceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	provinceModel := model.Province{}
	isThere, err := provinceModel.IsProvinceExistsById(provinceId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	"Not Found",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if isThere {
		cities, err := provinceModel.FindAllCityByProvinceId(provinceId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message: 	"Somethings wrong!",
				Errors: 	fmt.Sprintf("%s", err),
			})
			return
		}

		if len(cities) != 0 {
			c.JSON(http.StatusOK, response.ResponseCities {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the cities by provinceId",
				Cities:		cities,
			})
			return
		} else {
			c.JSON(http.StatusNotFound, response.ResponseGeneric {
				StatusCode:	http.StatusNotFound,
				Message:	"Can't found the cities by provinceId",
			})
			return
		}
	} else {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	"Not Found",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}
