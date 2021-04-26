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
		return false, err
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

	if err != nil {
		return &Province{}, err
	}

	currentId, err := result.LastInsertId()
	if err != nil {
		return &Province{}, err
	}

	p.ProvinceId = int(currentId)
	return p, nil
}

func (p *Province) FindProvinceById(provinceId int) (*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Province{}, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT province_id, province, created_at, updated_at FROM province WHERE province_id = ?", provinceId).
		Scan(&p.ProvinceId, &p.Province, &p.Audit.CreatedAt, &p.Audit.UpdatedAt)
	if err != nil {
		return &Province{}, err
	}

	if p  == (&Province{}) {
		return &Province{}, errors.New("The province can't found!")
	}

	return p, nil
}


func (p *Province) UpdateProvinceById(provinceId int) (*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Province{}, err
	}
	defer db.Close()

	result, err := db.Exec("UPDATE province SET province = ?, created_at = ?, updated_at = ? WHERE province_id = ?",
			p.Province,
			p.Audit.CreatedAt,
			p.Audit.UpdatedAt,
			provinceId)
	if err != nil {
		return &Province{},err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &Province{}, err
	}

	if rowsAffected != 1 {
		return &Province{}, errors.New("Somethings wrong!")
	}

	return p, nil
}

func (p *Province) DeleteProvinceById(provinceId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM province WHERE province_id = ?", provinceId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}

func (p *Province) FindAllProvince() ([]*Province, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*Province{}, err
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

func (p *Province) FindAllCityByProvinceId(provinceId int) ([]*City, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*City{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT c.city_id, c.province_id, c.city, c.created_at, c.updated_at FROM city c INNER JOIN province p ON c.province_id = p.province_id HAVING c.province_id = ?",provinceId)
	if err != nil {
		return []*City{}, err
	}
	defer rows.Close()

	result := []*City{}
	for rows.Next() {
		each := &City{}
		err := rows.Scan(&each.CityId, &each.ProvinceId, &each.City, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			// TODO: continue, if there are errors
			return []*City{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*City{}, err
	}

	return result, nil
}
