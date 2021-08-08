package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"net/http"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func GetPermisisonById(c *gin.Context) {
	permissionId := 0

	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	if permissionId <= 0 {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	permissionModel := model.Permission{}
	isThere, err := permissionModel.IsPermissionExistsById(permissionId)
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
			Message:	"The Permission is not exists",
		})
		return
	} else if isThere {
		currentPermission, err := permissionModel.FindPermissionById(permissionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if currentPermission != (&model.Permission{}) {
			c.JSON(http.StatusOK, response.ResponsePermision {
				StatusCode:	http.StatusOK,
				Message:	"Success to get the permission",
				Permission:	*currentPermission,
			})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseErrors {
		StatusCode:	http.StatusInternalServerError,
		Message:	"The server can't handle the request",
		Errors:		fmt.Sprintf("%s", err),
	})
	return
}

func CreatePermission(c *gin.Context) {
	requestPermission := model.Permission{}
	roleModel := model.Role{}

	err := c.Bind(&requestPermission)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid",
			Errors:		"Bad Request",
		})
		return
	}

	if requestPermission.RoleId <= 0 {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	isThereRole, err := roleModel.IsRoleExistsById(requestPermission.RoleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"The server can't handle the request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if !isThereRole {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The Role is not exists",
		})
		return
	} else if isThereRole {
		if len(strings.Trim(requestPermission.Permission, " ")) == 0 {
			c.JSON(http.StatusBadRequest, response.ResponseGeneric {
				StatusCode:	http.StatusBadRequest,
				Message:	"Permission name can't be emtpy",
			})
			return
		}

		requestPermission.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
		createdPermission, err := requestPermission.SavePermission()
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if createdPermission != (&model.Permission{}) {
			c.JSON(http.StatusOK, response.ResponsePermision {
				StatusCode:	http.StatusOK,
				Message:	"Success to create the permission",
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
