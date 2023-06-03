package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sales-backend/model"
	"sales-backend/response"
	"sales-backend/utility"

	"github.com/gin-gonic/gin"
)

func GetRoleById(c *gin.Context) {
	roleId := 0

	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	if roleId <= 0 {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	roleModel := model.Role{}
	isThere, err := roleModel.IsRoleExistsById(roleId)
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
			Message:    "The Role is not exists",
		})
		return
	} else if isThere {
		currentRole, err := roleModel.FindRoleById(roleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if currentRole != (&model.Role{}) {
			c.JSON(http.StatusOK, response.ResponseRole{
				StatusCode: http.StatusOK,
				Message:    "Success to get the role",
				Role:       *currentRole,
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

func CreateRole(c *gin.Context) {
	requestRole := model.Role{}

	err := c.Bind(&requestRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	if len(strings.Trim(requestRole.Role, " ")) == 0 {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "Role name can't be empty",
		})
		return
	}

	requestRole.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	createdRole, err := requestRole.SaveRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if createdRole != (&model.Role{}) {
		c.JSON(http.StatusOK, response.ResponseRole{
			StatusCode: http.StatusOK,
			Message:    "Success to create the role",
			Role:       *createdRole,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func UpdateRole(c *gin.Context) {
	requestRole := model.Role{}
	roleModel := model.Role{}
	roleId := 0

	err := c.Bind(&requestRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	roleId, err = strconv.Atoi(c.Param("roleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusOK,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	if roleId <= 0 {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	if roleId != requestRole.RoleId || (roleId <= 0 || requestRole.RoleId <= 0) {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Format!",
		})
		return
	}

	isThere, err := roleModel.IsRoleExistsById(roleId)
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
			Message:    "The Role is not exists",
		})
		return
	} else if isThere {
		currentRole, err := roleModel.FindRoleById(roleId)

		if len(strings.Trim(requestRole.Role, " ")) == 0 {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "Role name can't be empty",
			})
			return
		}

		if requestRole.Audit.CreatedAt != currentRole.Audit.CreatedAt {
			c.JSON(http.StatusBadRequest, response.ResponseGeneric{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid field createdAt",
			})
			return
		}

		// if the role already updated, should I check the field UpdatedAt?

		currentRole.RoleId = requestRole.RoleId
		currentRole.Role = requestRole.Role
		currentRole.Audit.CreatedAt = requestRole.Audit.CreatedAt
		timestamp := time.Now().Format("2006-01-02 15:05:03")
		currentRole.Audit.UpdatedAt = &timestamp

		updatedRole, err := currentRole.UpdateRoleById(roleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if updatedRole != (&model.Role{}) {
			c.JSON(http.StatusOK, response.ResponseRole{
				StatusCode: http.StatusOK,
				Message:    "Success to update the role",
				Role:       *updatedRole,
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

func DeleteRole(c *gin.Context) {
	roleId := 0

	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	if roleId <= 0 {
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid",
			Errors:     "Bad Request",
		})
		return
	}

	roleModel := model.Role{}
	isThere, err := roleModel.IsRoleExistsById(roleId)
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
			Message:    "The Role is not exist",
		})
		return
	} else if isThere {
		isSuccess, err := roleModel.DeleteRoleById(roleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		if !isSuccess {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusNotFound,
				Message:    "The Role is not exists",
			})
			return
		} else if isSuccess {
			c.JSON(http.StatusNotFound, response.ResponseGeneric{
				StatusCode: http.StatusOK,
				Message:    "Success to delete the role",
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

func GetRoles(c *gin.Context) {
	requestPage := c.DefaultQuery("page", "1")
	requestLimit := c.DefaultQuery("limit", "10")
	roleModel := model.Role{}

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

	numberRecords, err := roleModel.GetNumberRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     "Internal Error",
		})
		return
	}

	nextPage, prevPage, totalPages := utility.GetPaginateURL([]string{"roles"}, &page, &limit, numberRecords)
	offset := limit * (page - 1)

	roles, err := roleModel.FindAllRole(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "Somethings wrong!",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if len(roles) != 0 {
		c.JSON(http.StatusOK, response.ResponseRoles{
			StatusCode: http.StatusOK,
			Message:    "Success to get the roles",
			Roles:      roles,
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
