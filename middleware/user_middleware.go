package middleware

import (
	"log"
	"fmt"
	"regexp"
	"io/ioutil"
	"bytes"
	"net/http"
	"net/mail"
	"strings"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRegisterModel := model.UserRegister{}
		requestUserRegister := model.UserRegister{}

		buf, _ := ioutil.ReadAll(c.Request.Body)
		currentCheck := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = currentCheck

		err := c.Bind(&requestUserRegister)
		if err != nil {
			log.Print("%s", err)
			c.JSON(http.StatusBadRequest, response.ResponseErrors {
				StatusCode:	http.StatusBadRequest,
				Message:	"Invalid Request",
				Errors:		"Bad Request",
			})
			return
		}

		// Create another 'request which used by another handler
		requestUserRegisterAgain := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = requestUserRegisterAgain

		isThereInvalid := false
		listInvalid := make(map[string]string)

		// log.Printf("%v", requestUserRegister)

		// validate the RoleId is exists.
		// This part will be repalce with associated struct
		if requestUserRegister.RoleId <= 0 {
			isThereInvalid = true
			listInvalid["err_role_id"] = "The RoleId is invalid"
		} else {
			isThereRoleId, err := userRegisterModel.IsRoleExistsById(requestUserRegister.RoleId)
			if err != nil {
				log.Printf("%s", err)
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"The server can't handle the request",
					Errors:		fmt.Sprintf("%s", err),
				})
				return
			}

			if !isThereRoleId {
				isThereInvalid = true
				listInvalid["err_role_id"] = "The RoleId is not exists."
			}
		}

		// validate the ProvinceId is exists
		if requestUserRegister.ProvinceId <= 0 {
			isThereInvalid = true
			listInvalid["err_province_id"] = "The ProvinceId is invalid"
		} else {
			provinceModel := model.Province{}
			var provinceId int = requestUserRegister.ProvinceId
			isThereProvinceId, err := provinceModel.IsProvinceExistsById(provinceId)
			if err != nil {
				log.Print("%s", err)
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"The server can't handle the request",
					Errors:		fmt.Sprintf("%s", err),
				})
				return
			}

			if !isThereProvinceId {
				isThereInvalid = true
				listInvalid["err_province_id"] = "The Provinceid is not exists."
			}
		}

		// validate the CityId is exisst
		if requestUserRegister.CityId <= 0 {
			isThereInvalid = true
			listInvalid["err_city_id"] = "The CityId is invalid"
		} else {
			cityModel := model.City{}
			var cityId int = requestUserRegister.CityId
			isThereCityId, err := cityModel.IsCityExistsById(cityId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"The server can't handle the request",
					Errors:		fmt.Sprintf("%s", err),
				})
				return
			}

			if !isThereCityId {
				isThereInvalid = true
				listInvalid["err_city_id"] = "The CityId is not exists."
			}
		}

		// validate the SubDistrictId is exists
		if requestUserRegister.SubDistrictId <= 0 {
			isThereInvalid = true
			listInvalid["err_sub_district_id"] = "The SubDistrictId is invalid"
		} else {
			subDistrictModel := model.SubDistrict{}
			isThereSubDistrictId, err := subDistrictModel.IsSubDistrictExistsById(requestUserRegister.SubDistrictId)
			if err != nil {
				log.Printf("%s", err)
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"The server can't handle the request",
					Errors:		fmt.Sprintf("%s", err),
				})
				return
			}

			if !isThereSubDistrictId {
				isThereInvalid = true
				listInvalid["err_sub_district_id"] = "The SubDistrictId is not exists"
			}
		}

		// validate the Email is not used by another user
		if _, err = mail.ParseAddress(requestUserRegister.Email); err != nil {
			isThereInvalid = true
			listInvalid["err_email"] = "The Email is invalid"
		} else {
			isThereEmail, err := userRegisterModel.IsUserExistsByEmail(requestUserRegister.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"The server can't handle the request",
					Errors:		fmt.Sprintf("%s", err),
				})
				return
			}

			if isThereEmail {
				isThereInvalid = true
				listInvalid["err_email"] = "The Email is already used."
			}
		}

		// validate the FirstName field
		if len(strings.Trim(requestUserRegister.FirstName, " ")) == 0 {
			isThereInvalid = true
			listInvalid["err_first_name"] = "The FirstName can't be empty"
		}

		// validate the Gender field
		if len(strings.Trim(requestUserRegister.Gender, " ")) == 0 {
			isThereInvalid = true
			listInvalid["err_gender"] = "The Gender can't be empty"
		} else if requestUserRegister.Gender != "Laki - laki" || requestUserRegister.Gender == "Perempuan" {
			isThereInvalid = true
			listInvalid["err_gender"] = "The Gender is invalid"
		}

		// validate the Address field
		if len(strings.Trim(requestUserRegister.Address, "" )) == 0 {
			isThereInvalid = true
			listInvalid["err_address"] = "The Address can't be empty"
		}

		// validate the PhoneNumber field
		re := regexp.MustCompile("(0|\\+62|062|62)[0-9]+$")
		if !re.MatchString(requestUserRegister.PhoneNumber) {
			isThereInvalid = true
			listInvalid["err_phone_number"] = "The PhoneNumber is invalid"
		} else {
			isTherePhoneNumber, err := userRegisterModel.IsPhoneNumberExists(requestUserRegister.PhoneNumber)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.ResponseErrors {
					StatusCode:	http.StatusInternalServerError,
					Message:	"The server can't handle the request",
					Errors:		fmt.Sprintf("%s", err),
				})
				return
			}

			if isTherePhoneNumber {
				isThereInvalid = true
				listInvalid["err_phone_number"] = "The PhoneNumber is already used."
			}
		}

		if isThereInvalid {
			c.JSON(http.StatusBadRequest, response.ResponseInvalids {
				StatusCode:	http.StatusBadRequest,
				Message:	"Make sure the fields is valid",
				Invalid:	listInvalid,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
