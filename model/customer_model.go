package model

import (
	"errors"

	"sales-backend/utility"
)

const (
	CUSTOMER_ROLE = "customer"
)

type Customer struct {
	CustomerId  string  `json:"customerId"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Gender      string  `json:"gender"`
	Address     Address `json:"address"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Audit       Audit   `json:"audit"`
}

func (c *Customer) IsCustomerExsistById(customerId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	checkUserRole := 0
	err = tx.QueryRow("SELECT COUNT(user_roles_id) FROM users_roles WHERE user_id = ?", customerId).Scan(&checkUserRole)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	// the user not have customer role
	if checkUserRole == 0 {
		tx.Rollback()
		return false, nil
	}

	roleId := 0
	err = tx.QueryRow("SELECT role_id FROM users_roles WHERE user_id = ?", customerId).Scan(&roleId)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if roleId == 0 {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	checkRole := 0
	role := ""
	err = tx.QueryRow("SELECT COUNT(role_id), role_name FROM roles WHERE role_id = ?", roleId).Scan(&checkRole, &role)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if checkRole != 1 || role != CUSTOMER_ROLE {
		tx.Rollback()
		return false, nil
	}

	if checkRole == 1 && role == CUSTOMER_ROLE {
		return true, nil
	}

	return false, errors.New("Somethings wrong!")
}
