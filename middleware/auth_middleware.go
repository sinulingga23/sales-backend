package middleware

import (
	"fmt"
	"net/http"

	"sales-backend/auth"
	"sales-backend/model"
	"sales-backend/response"

	"github.com/gin-gonic/gin"
)

func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isValidToken, err := auth.IsValidToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ResponseGeneric{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token is invalid",
			})
			c.Abort()
			return
		}

		if !isValidToken {
			c.JSON(http.StatusUnauthorized, response.ResponseGeneric{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token is invalid.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ValidateAdminMiddleware() gin.HandlerFunc {
	role := "admin"
	return func(c *gin.Context) {
		roleId, err := auth.ExtractTokenRoleId(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ResponseGeneric{
				StatusCode: http.StatusUnauthorized,
				Message:    fmt.Sprintf("%s", err),
			})
			c.Abort()
			return
		}

		roleModel := model.Role{}
		currentRole, err := roleModel.GetRoleById(roleId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors{
				StatusCode: http.StatusInternalServerError,
				Message:    "The server can't handle the request",
				Errors:     fmt.Sprintf("%s", err),
			})
			c.Abort()
			return
		}

		if currentRole != role {
			c.JSON(http.StatusUnauthorized, response.ResponseGeneric{
				StatusCode: http.StatusUnauthorized,
				Message:    "The role is invalid",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
