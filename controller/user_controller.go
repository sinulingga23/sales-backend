package controller

import (
	"fmt"
	"log"
	"net/http"
	_ "strconv"
	"strings"
	"time"

	"github.com/sinulingga23/sales-backend/model"
	"github.com/sinulingga23/sales-backend/response"
	_ "github.com/sinulingga23/sales-backend/utility"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	log.Printf("%s", userId)

	if len(strings.Trim(userId, " ")) == 0 {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric{
			StatusCode: http.StatusBadRequest,
			Message:    "UserId can't be empty",
		})
		return
	}

	userRegisterModel := model.UserRegister{}
	isThere, err := userRegisterModel.IsUserExistsById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		log.Printf("line 36 gan: %v\n", err)
		return
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric{
			StatusCode: http.StatusNotFound,
			Message:    "The user is not exist.",
		})
		return
	} else if isThere {
		currentUser, err := userRegisterModel.FindUserById(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "Somethings wrong!",
				Errors:     fmt.Sprintf("%s", err),
			})
			log.Printf("line 54: %s\n", err)
			return
		}

		if currentUser != (&model.User{}) {
			c.JSON(http.StatusOK, response.ResponseUser{
				StatusCode: http.StatusOK,
				Message:    "Success to get the user",
				User:       *currentUser,
			})
			log.Printf("line 64: %s\n", err)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}

func CreateUser(c *gin.Context) {
	requestUserRegister := model.UserRegister{}
	err := c.Bind(&requestUserRegister)
	if err != nil {
		log.Printf("line 81 %s", err)
		c.JSON(http.StatusBadRequest, response.ResponseErrors{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
			Errors:     "Bad Request",
		})
		return
	}

	requestUserRegister.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	createdUser, err := requestUserRegister.SaveUser()
	if err != nil {
		log.Printf("line 93 %s", err)
		c.JSON(http.StatusInternalServerError, response.ResponseErrors{
			StatusCode: http.StatusInternalServerError,
			Message:    "The server can't handle the request",
			Errors:     fmt.Sprintf("%s", err),
		})
		return
	}

	if createdUser != (&model.User{}) {
		c.JSON(http.StatusOK, response.ResponseUser{
			StatusCode: http.StatusOK,
			Message:    "Success to create the user",
			User:       *createdUser,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric{
		StatusCode: http.StatusInternalServerError,
		Message:    "Somethings wrong!",
	})
	return
}
