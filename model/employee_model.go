package model

import (
	"errors"

	"sales-backend/utility"
)

const (
	EMPLOYEE_ROLE = "employee"
)

type Employee struct {
	EmployeeId  string  `json:"employeeId"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Gender      string  `json:"gender"`
	Address     Address `json:"address"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Audit       Audit   `json:"audit"`
}

func (e *Employee) IsEmployeeExistsById(employeeId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	checkUserRole := 0
	err = tx.QueryRow("SELECT COUNT(user_roles_id) FROM users_roles WHERE user_id = ?", employeeId).Scan(&checkUserRole)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if checkUserRole == 0 {
		tx.Rollback()
		return false, nil
	}

	roleId := 0
	err = tx.QueryRow("SELECT role_id FROM users_roles WHERE user_id = ?", employeeId).Scan(&roleId)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if roleId == 0 {
		return false, nil
	}

	checkRole := 0
	role := ""
	err = tx.QueryRow("SELECT COUNT(role_id), role_name FROM roles WHERE role_id = ?", roleId).Scan(&checkRole, &role)
	if err != nil {
		tx.Rollback()
		return false, errors.New("Somethings wrong!")
	}

	if checkRole != 1 || role != EMPLOYEE_ROLE {
		tx.Rollback()
		return false, nil
	}

	if checkRole == 1 && role == EMPLOYEE_ROLE {
		return true, nil
	}

	return false, errors.New("Somethings wrong!")
}
