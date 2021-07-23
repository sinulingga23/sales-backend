package auth

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func CreateToken(email string, roleId int) (string, error) {
	if len(os.Getenv("JWT_SECRET_KEY")) == 0 {
		return "", errors.New("Somethings wrong!")
	}

	claims := struct {
		jwt.StandardClaims
		Email	string
		RoleId	int
	}{
		jwt.StandardClaims{
			ExpiresAt: 	3600 * 24 * 30,
			Issuer:		email,
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
	if bearerToken = ExtractToken(r); bearerToken != "" {
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
	if bearerToken = ExtractToken(r); bearerToken != "" {
		return "", errors.New("Token is invalid")
	}
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]);
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

func ExtractTokenRoleId(r *http.Request) (int, error) {
	var bearerToken string
	if bearerToken = ExtractToken(r); bearerToken != "" {
		return 0, errors.New("Token is invalid")
	}
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]);
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		roleId, err := strconv.ParseInt(fmt.Sprintf("%0.f", claims["role_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(roleId), nil
	}

	return 0, errors.New("Somethings wrong!")
}


