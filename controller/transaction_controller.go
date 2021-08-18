package controller

import (
	"fmt"
	"time"
	"strings"
	"net/http"

	"sales-backend/model"
	"sales-backend/response"
	"github.com/gin-gonic/gin"
)

func GetTransactionById(c *gin.Context) {
	transactionId := c.Param("transactionId")

	if len(strings.Trim(transactionId, " ")) == 0 {
		c.JSON(http.StatusBadRequest, response.ResponseGeneric {
			StatusCode:	http.StatusBadRequest,
			Message:	"TransactionId can't be empty",
		})
		return
	}

	transactionModel := model.Transaction{}
	isThere, err := transactionModel.IsTransactionExistsById(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle the request",
			Errors:		fmt.Sprintf("%s", err),
		})
	}

	if !isThere {
		c.JSON(http.StatusNotFound, response.ResponseGeneric {
			StatusCode:	http.StatusNotFound,
			Message:	"The transaction is not exists",
		})
		return
	} else if isThere {
		currentTransaction, err := transactionModel.FindTransactionById(transactionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
			})
			return
		}

		if currentTransaction != (&model.Transaction{}) {
			c.JSON(http.StatusOK, struct {
				StatusCode	int 			`json:"statusCode"`
				Message		string			`json:"message"`
				Transaction 	model.Transaction	`json:"transaction"`
			}{http.StatusOK, "Success to get the transaction", *currentTransaction})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}

func CreateTransaction(c *gin.Context) {
	requestTransaction := model.Transaction{}

	err := c.Bind(&requestTransaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseErrors {
			StatusCode:	http.StatusBadRequest,
			Message:	"Invalid Request",
			Errors:		"Bad Request",
		})
		return
	}

	isThereInvalid := false
	listInvalid := make(map[string]string)

	if len(strings.Trim(requestTransaction.CustomerId, " ")) == 0 {
		isThereInvalid = true
		listInvalid["err_customer_id"]  = "CustomerId can't be empty"
	}

	if len(strings.Trim(requestTransaction.EmployeeId, " ")) == 0 {
		isThereInvalid = true
		listInvalid["err_employee_id"] = "EmployeeId can't be empty"
	}

	if len(strings.Trim(requestTransaction.CustomerId, " ")) != 0 {
		customerModel := model.Customer{}
		isThereCustomer, err := customerModel.IsCustomerExsistById(requestTransaction.CustomerId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if !isThereCustomer {
			isThereInvalid = true
			listInvalid["err_customer_id"] = "The Customer is not exists"
		}
	}

	if len(strings.Trim(requestTransaction.EmployeeId, " ")) != 0 {
		employeeModel := model.Employee{}
		isThereEmployee, err := employeeModel.IsEmployeeExistsById(requestTransaction.EmployeeId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ResponseErrors {
				StatusCode:	http.StatusInternalServerError,
				Message:	"The server can't handle the request",
				Errors:		fmt.Sprintf("%s", err),
			})
			return
		}

		if !isThereEmployee {
			isThereInvalid = true
			listInvalid["err_employee_id"]  = "The Employee is not exists"
		}
	}

	if isThereInvalid {
		// 400 or 404?
		c.JSON(http.StatusBadRequest, response.ResponseInvalids {
			StatusCode:	http.StatusBadRequest,
			Message:	"Make sure the field is valid",
			Invalid:	listInvalid,
		})
		return
	}


	requestTransaction.Date = time.Now().Format("2006-01-02")
	requestTransaction.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	createdTransaction, err := requestTransaction.SaveTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseErrors {
			StatusCode:	http.StatusInternalServerError,
			Message:	"The server can't handle the request",
			Errors:		fmt.Sprintf("%s", err),
		})
		return
	}

	if createdTransaction != (&model.Transaction{}) {
		c.JSON(http.StatusOK, struct {
			StatusCode	int 			`json:"statusCode"`
			Message		string			`json:"message"`
			Transaction 	model.Transaction	`json:"transaction"`
		}{http.StatusOK, "Success to create the transaction", *createdTransaction})
		return
	}

	c.JSON(http.StatusInternalServerError, response.ResponseGeneric {
		StatusCode:	http.StatusInternalServerError,
		Message:	"Somethings wrong!",
	})
	return
}
