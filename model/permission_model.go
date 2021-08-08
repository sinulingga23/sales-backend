package model

import (
	"errors"

	"sales-backend/utility"
)

type Permission struct {
	PermissionId	int 	`json:"permissionId"`
	RoleId		int 	`json:"roleId"`
	Permission	string	`json:"permission"`
	Audit		Audit	`json:"audit"`
}

func (p *Permission) IsPermissionExistsById(permissionId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(permission_id) FROM permissions WHERE permission_id = ?", permissionId).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (p *Permission) SavePermission() (*Permission, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO permissions (role_id, permission_name, created_at) VALUES (?, ?, ?)",
		p.RoleId,
		p.Permission,
		p.Audit.CreatedAt)
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	p.PermissionId = int(lastInsertId)
	return p, nil
}

func (p *Permission) FindPermissionById(permissionId int) (*Permission, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	err = db.QueryRow("SELECT permission_id, role_id, permission_name, created_at, updated_at FROM permissions WHERE permission_id = ?", permissionId).
		Scan(&p.PermissionId, &p.RoleId, &p.Permission, &p.Audit.CreatedAt, &p.Audit.UpdatedAt)
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	if p == (&Permission{}) {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	return p, nil
}

func (p *Permission) UpdatePermissionById(permissionId int) (*Permission, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE permissions SET role_id = ?, permission_name = ?, created_at = ?, updated_at = ? WHERE permision_id = ?",
		p.RoleId,
		p.Permission,
		p.Audit.CreatedAt,
		p.Audit.UpdatedAt,
		permissionId)
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return &Permission{}, errors.New("Somethings wrong!")
	}

	p.PermissionId = permissionId
	return p, nil
}

func (p *Permission) DeletePermissionById(permissionId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM permissions WHERE permission_id = ?", permissionId)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}

func (p *Permission) FindAllPermission(limit int, offset int) ([]*Permission, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*Permission{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT permission_id, role_id, permission_name, created_at, updated_at FROM permission LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return []*Permission{}, errors.New("Somethings wrong!")
	}

	result := []*Permission{}
	for rows.Next() {
		each := &Permission{}

		err = rows.Scan(&each.PermissionId, &each.RoleId, &each.Permission, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*Permission{}, errors.New("Somethings wrong!")
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*Permission{}, errors.New("Somethings wrong!")
	}

	return result, nil
}
