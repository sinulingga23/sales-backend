package model

import (
	"errors"

	"sales-backend/utility"
)

type Role struct {
	RoleId	int 	`json:"roleId"`
	Role	string 	`json:"role"`
	Audit	Audit	`json:"audit"`
}

func (r *Role) IsRoleExistsById(roleId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}

	check := 0
	err = db.QueryRow("SELECT COUNT(role_id) FROM roles WHERE role_id = ?", roleId).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (r *Role) GetRoleById(roleId int) (string, error) {
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

func (r *Role) FindRoleById(roleId int) (*Role, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Role{}, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT role_id, role_name, created_at, updated_at FROM roles WHERE role_id = ?", roleId).
		Scan(&r.RoleId, &r.Role, &r.Audit.CreatedAt, &r.Audit.UpdatedAt)
	if err != nil {
		return &Role{}, err
	}

	if r == (&Role{}) {
		return &Role{}, err
	}

	return r, nil
}

func (r *Role) SaveRole() (*Role, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Role{}, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO roles (role_name, created_at) VALUES (?, ?)", r.Role, r.Audit.CreatedAt)
	if err != nil {
		return &Role{}, err
	}

	currentId, err := result.LastInsertId()
	if err != nil {
		return &Role{}, err
	}

	r.RoleId = int(currentId)
	return r, nil
}

func (r *Role) UpdateRoleById(roleId int) (*Role, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Role{}, err
	}
	defer db.Close()

	result, err := db.Exec("UPDATE roles SET role_name = ?, created_at = ?, updated_at = ? WHERE role_id = ?",
		r.Role,
		r.Audit.CreatedAt,
		r.Audit.UpdatedAt,
		roleId)
	if err != nil {
		return &Role{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &Role{}, err
	}

	if rowsAffected != 1 {
		return &Role{}, errors.New("Somethings wrong!")
	}

	return r, nil
}

func (r *Role) DeleteRoleById(roleId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM roles WHERE role_id = ?", roleId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected != 1 {
		return false, err
	}

	return true, nil
}

func (r *Role) FindAllRole(limit int, offset int) ([]*Role, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*Role{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT role_id, role_name, created_at, updated_at FROM roles LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return []*Role{}, err
	}
	defer rows.Close()

	result := []*Role{}
	for rows.Next() {
		each := &Role{}
		err = rows.Scan(&each.RoleId, &each.Role, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*Role{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*Role{}, err
	}

	return result, nil
}

func (r *Role) GetNumberRecords() (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(role_id) FROM roles").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}
