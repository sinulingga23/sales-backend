package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: Create function ExtractElapsedTime()

func CreateToken(email string, roleId int) (string, error) {
	if len(os.Getenv("JWT_SECRET_KEY")) == 0 {
		return "", errors.New("Somethings wrong!")
	}

	claims := struct {
		jwt.StandardClaims
		Email  string
		RoleId int
	}{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    email,
		},
		email,
		roleId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var bearerToken string
	var err error
	if bearerToken, err = token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY"))); err != nil {
		return "", err
	}

	return bearerToken, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func IsValidToken(r *http.Request) (bool, error) {
	var bearerToken string
	if bearerToken = ExtractToken(r); bearerToken == "" {
		return false, errors.New("Token is invalid")
	}
	_, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractTokenEmail(r *http.Request) (string, error) {
	var bearerToken string
	if bearerToken = ExtractToken(r); bearerToken == "" {
		return "", errors.New("Token is invalid")
	}
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["email"].(string), nil
	}

	return "", errors.New("Somethings wrong!")
}

func ExtractTokenRoleId(r *http.Request) (string, error) {
	var bearerToken string
	if bearerToken = ExtractToken(r); bearerToken == "" {
		return "", errors.New("Token is invalid")
	}
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		roleId := claims["RoleId"]
		return roleId.(string), nil
	}

	return "", errors.New("Somethings wrong!")
}
