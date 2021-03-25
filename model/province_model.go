package model

import (
	"fmt"
	"errors"
	"time"

	"sales-backend/utility"
)

type Province struct {
	ProvinceId	int	`json:"provinceId"`
	Province   	string	`json:"province"`
	Audit		Audit	`json:"audit"`
}


func (p *Province) IsProvinceExists(provinceId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	// Make sure the important fields is not empty
	if provinceId == "" {
		return false, errors.New("ProvinceId can't be empty")
	}

	check := 0
	err = db.QueryRow("SELECT COUNT(province_id) FROM province WHERE province_id = ?", provinceId).Scan(&check)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if check != 0 {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("Can't find province with id: %s", provinceId))
}

func (p *Province) SaveProvince() (*Province, error) {
	// Make sure the important fields is not empty
	if p.Province == "" {
		return &Province{}, errors.New("Province name can't be empty")
	}

	// If the CreatedAt field is empty, then set the field using the current time
	if p.Audit.CreatedAt == "" {
		p.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	}

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


// func (p *Province) SaveProvince(*Province, errors) {
// 	db, err := utility.ConnectDB()
// 	if err != nil {
// 		return false, err
// 	}
// 	defer db.Close()

// 	if
// }
