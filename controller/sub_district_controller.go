package controller


import (
	"log"
	"fmt"
	"time"
	"strings"
	"strconv"
	"net/http"

	"sales-backend/model"
	"sales-backend/response"
	"sales-backend/utility"
	"github.com/gin-gonic/gin"
)

func GetSubDistrictById(c *gin.Context) {
	subDistrictId := 0

	subDistrictId, err := strconv.Atoi(c.Param("subDistrictId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid",
			Errors:		"Bad Request",
		})
		return
	}

	subDistrictModel := model.SubDistrict{}
	isThere, err := subDistrictModel.IsSubDistrictExistsById(subDistrictId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The Sub-District is not exists",
		})
		return
	} else if isThere {
		currentSubDistrict, err := subDistrictModel.FindSubDistrictById(subDistrictId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"Somethings wrong!",
				Errors:		fmt.Sprintf("%s", err),
			})
			log.Printf("%s", err)
			return
		}

		if currentSubDistrict != (&model.SubDistrict{}) {
			c.JSON(http.StatusOK, response.ResponseSubDistrict {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the sub-district",
				SubDistrict:	*currentSubDistrict,
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

func CreateSubDistrict(c *gin.Context) {
	requestSubDistrict := model.SubDistrict{}
	cityModel := model.City{}

	err := c.Bind(&requestSubDistrict)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message: 	"Invalid Request",
			Errors: 	"Bad Request",
		})
		log.Printf("%s", err)
		return
	}

	if requestSubDistrict.CityId < 0 {
		c.JSON(http.StatusNotFound, response.ResponseErrors {
			StatusCode:	http.StatusNotFound,
			Message:	"Somethings wrong!",
			Errors:		fmt.Sprintf("The City with id %d is not exists.", requestSubDistrict.CityId),
		})
		return
	}
	isThereCity, err := cityModel.IsCityExistsById(requestSubDistrict.CityId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThereCity {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The City is not exists",
		})
		return
	} else if isThereCity {
		if strings.Trim(requestSubDistrict.SubDistrict, " ") == "" {
			c.JSON(http.StatusBadRequest, response.ResponseGeneric {
				StatusCode:	http.StatusBadRequest,
				Message: 	"SubDistrict name can't be empty",
			})
			log.Printf("%s", err)
			return
		} else  {
			requestSubDistrict.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
			createdSubDistrict, err := requestSubDistrict.SaveSubDistrict()
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"Somethings wrong!",
					Errors:		fmt.Sprintf("%v", err),
				})
				log.Printf("%s", err)
				return
			}

			if createdSubDistrict != (&model.SubDistrict{}) {
				c.JSON(http.StatusOK, response.ResponseSubDistrict {
					StatusCode:	http.StatusOK,
					Message:	"Success to create a subDistrict.",
					SubDistrict:	*createdSubDistrict,
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

func UpdateSubDistrictById(c *gin.Context) {
	subDistrictId := 0
	requestSubDistrict := model.SubDistrict{}
	cityModel := model.City{}

	err := c.Bind(&requestSubDistrict)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Somethings wrong!",
		})
		log.Print("%s", err)
		return
	}

	subDistrictId, err = strconv.Atoi(c.Param("subDistrictId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Somethings wrong!",
		})
		log.Printf("%s", err)
		return
	}

	if subDistrictId != requestSubDistrict.SubDistrictId || (subDistrictId <= 0 || requestSubDistrict.SubDistrictId <= 0){
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid format!",
		})
		return
	}

	if strings.Trim(requestSubDistrict.SubDistrict, " ") == "" {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"SubDistrict can't be empty",
		})
		return
	}

	isThereCity, err := cityModel.IsCityExistsById(requestSubDistrict.CityId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThereCity {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The City is not exists",
		})
		return
	} else if isThereCity {
		subDistrictModel := model.SubDistrict{}
		isThereSubDistrict, err := subDistrictModel.IsSubDistrictExistsById(requestSubDistrict.SubDistrictId)
		if err != nil {
			log.Printf("%s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if !isThereSubDistrict {
			c.JSON(http.StatusNotFound, response.ResponseGeneric {
				StatusCode:	http.StatusNotFound,
				Message:	"The Sub-District is not exists",
			})
			return
		} else if isThereSubDistrict {
			currentSubDistrict, err := subDistrictModel.FindSubDistrictById(subDistrictId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
					StatusCode:	http.StatusInternalServerError,
					Message:	fmt.Sprintf("%s", err),
				})
				log.Printf("%s", err)
				return
			}

			if currentSubDistrict.Audit.CreatedAt != requestSubDistrict.Audit.CreatedAt {
				c.JSON(http.StatusBadRequest, response.ResponseGeneric {
					StatusCode:	http.StatusBadRequest,
					Message:	"Invalid field createdAt",
				})
				return
			}

			currentSubDistrict.SubDistrictId = requestSubDistrict.SubDistrictId
			currentSubDistrict.CityId = requestSubDistrict.CityId
			currentSubDistrict.SubDistrict = requestSubDistrict.SubDistrict
			currentSubDistrict.Audit.CreatedAt = requestSubDistrict.Audit.CreatedAt
			timestamp := time.Now().Format("2006-01-02 15:05:03")
			currentSubDistrict.Audit.UpdatedAt = &timestamp


			updatedSubDistrict, err := currentSubDistrict.UpdateSubDistrictById(subDistrictId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
					StatusCode:	http.StatusInternalServerError,
					Message:	"Somethings wrong!",
				})
				log.Printf("%s", err)
				return
			}

			if updatedSubDistrict != (&model.SubDistrict{}) {
				c.JSON(http.StatusOK, response.ResponseSubDistrict {
					StatusCode:	http.StatusOK,
					Message:	"Success to update the subDistrict",
					SubDistrict:	*updatedSubDistrict,
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

func DeleteSubDistrictById(c *gin.Context) {
	subDistrictId := 0

	subDistrictId, err := strconv.Atoi(c.Param("subDistrictId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid",
			Errors:		"Bad Request",
		})
		log.Printf("%s", err)
		return
	}

	cityModel := model.SubDistrict{}
	isThere, err := cityModel.IsSubDistrictExistsById(subDistrictId)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The Sub-District is not exists",
		})
		return
	} else if isThere {
		isDeleted, err := cityModel.DeleteSubDistrictById(subDistrictId)
		if err != nil {
			log.Printf("%s", err)
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if !isDeleted {
			c.JSON(http.StatusNotFound, response.ResponseGeneric {
				StatusCode:	http.StatusNotFound,
				Message:	"The Sub-District is not exists",
			})
			return
		} else if isDeleted {
			c.JSON(http.StatusOK, response.ResponseGeneric {
				StatusCode:	http.StatusOK,
				Message:	"Success to delete the Sub-districts.",
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

func GetSubDistricts(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	subDistrictModel := model.SubDistrict{}

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

	numberRecords, err := subDistrictModel.GetNumberRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"Somethings wrong!",
			Errors:		"Internal Error",
		})
		return
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"sub-districts"}, &page, &limit, numberRecords)
	offset := limit * (page-1)

	subDistricts, err := subDistrictModel.FindAllSubDistrict(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message: 	"Somethings wrong!",
			Errors: 	fmt.Sprintf("%s", err),
		})
		return
	}

	if len(subDistricts) != 0 {
		c.JSON(http.StatusOK, response.ResponseSubDistricts {
			StatusCode:	http.StatusOK,
			Message:	"Success to get the sub-districts",
			SubDistricts:	subDistricts,
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
			Message:	"The sub-districts is empty",
		})
		return
	}

	c.JSON(http.StatusNotFound, response.ResponseGeneric {
		StatusCode:	http.StatusNotFound,
		Message:	"The cities is empty",
	})
	return
}
