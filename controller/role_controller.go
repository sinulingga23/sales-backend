package controller

import (
	"fmt"
	"net/http"
	"srtconv"

	"sales-backend/model"
	"github.com/gin-gonic/gin"
)

func GetRoleById(c *gin.Context) {
	roleId := 0

	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid",
			Errors:		"Bad Request",
		})
		return
	}

	roleModel := model.Role{}
	isThere, err := roleModel.IsRoleExistsById(roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle the request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The Role is not exists",
		})
		return
	} else if isThere {
		currentRole, err := roleModel.GetRoleById(roleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if currentRole != (model.Role{}) {
			c.JSON(http.StatusOK, response.ResponseRole {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the role",
				Role:		*currentRole
			})
		}
	}
	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}
