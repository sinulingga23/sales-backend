package model

import (
	"sales-backend/utility"
)

type Role struct {
}

func (r Role) GetRoleById(roleId int) (string, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	role := ""
	err = db.QueryRow("SELECT role_name FROM roles WHERE role_id = ?", roleId).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}
