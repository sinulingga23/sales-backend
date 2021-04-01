package model

import (
	"fmt"
	"errors"

	"sales-backend/utility"
)

type Province struct {
	ProvinceId	int	`json:"provinceId"`
	Province   	string	`json:"province"`
	Audit		Audit	`json:"audit"`
}


func (p *Province) IsProvinceExistsById(provinceId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(province_id) FROM province WHERE province_id = ?", provinceId).Scan(&check)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if check != 1 {
		return false, errors.New(fmt.Sprintf("The province with id %d is not exists.", provinceId))
	}
	return true, nil
}

func (p *Province) SaveProvince() (*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Province{}, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO province (province, created_at) VALUES (?, ?)",
		p.Province,
		p.Audit.CreatedAt)

	if err != nil  || result != nil {
		return &Province{}, errors.New("Somethings wrong!")
	}

	currentId, err := result.LastInsertId()
	if err != nil {
		return &Province{}, errors.New("Somethings wrong!")
	}

	p.ProvinceId = int(currentId)
	return p, nil
}

func (p *Province) FindProvinceById(provinceId int) (*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Province{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	err = db.QueryRow("SELECT province_id, province, created_at, updated_at FROM province WHERE province_id = ?", provinceId).
		Scan(&p.ProvinceId, &p.Province, &p.Audit.CreatedAt, &p.Audit.UpdatedAt)
	if err != nil {
		return &Province{}, errors.New("Somethings wrong!")
	}

	if p  == (&Province{}) {
		return &Province{}, errors.New("The province can't found!")
	}

	return p, nil
}


func (p *Province) UpdateProvinceById(provinceId int) (*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil || p.ProvinceId != provinceId {
		return &Province{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE province SET province = ?, created_at = ?, updated_at = ? WHERE province_id = ?",
			p.Province,
			p.Audit.CreatedAt,
			p.Audit.UpdatedAt,
			provinceId)
	if err != nil {
		return &Province{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &Province{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return &Province{}, errors.New("Somethings wrong!")
	}

	return p, nil
}

func (p *Province) DeleteProvinceById(provinceId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("DELETE province WHERE province_id = ?", provinceId)
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

func (p *Province) FindAllProvince() ([]*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*Province{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT province_id, province, created_at, updated_at FROM province")
	if err != nil {
		return []*Province{}, errors.New("Somethings wrong!")
	}
	defer rows.Close()

	result := []*Province{}
	for rows.Next() {
		each := &Province{}

		err = rows.Scan(&each.ProvinceId, &each.Province, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*Province{}, errors.New("Somethings wrong!")
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*Province{}, errors.New("Somethings wrong!")
	}

	return result, nil
}
