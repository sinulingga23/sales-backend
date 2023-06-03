package controller

import (
	"fmt"
	"net/http"

	"github.com/sinulingga23/sales-backend/auth"
	"github.com/sinulingga23/sales-backend/model"
	"github.com/sinulingga23/sales-backend/response"

	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	requestEmail, requestPassword, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, response.ResponseGeneric{
			StatusCode: http.StatusUnauthorized,
			Message:    "Login is not valid",
		})
		return
	}

	loginModel := model.Login{}
	isLoginValid, err := loginModel.Login(requestEmail, requestPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ResponseGeneric{
			StatusCode: http.StatusUnauthorized,
			Message:    "The email/password is wrong",
		})
		return
	}

	if isLoginValid {
		userRegisterModel := model.UserRegister{}
		roleId, err := userRegisterModel.GetRoleIdByEmail(requestEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		bearerToken, err := auth.CreateToken(requestEmail, roleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			return
		}

		c.JSON(http.StatusOK, response.ResonseJWTToken{
			StatusCode: http.StatusOK,
			Token:      bearerToken,
		})
		return
	}
}
