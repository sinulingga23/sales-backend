package model

import (
	"errors"
	"fmt"
	"log"

	"sales-backend/utility"
)

type Address struct {
	AddressId     int    `json:"addressId"`
	ProvinceId    int    `json:"provinceId"`
	CityId        int    `json:"cityId"`
	SubDistrictId int    `json:"subDistrictId"`
	Address       string `json:"address"`
	Audit         Audit  `json:"audit"`
}

func (a *Address) IsAddressExistsById(addressId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return false, err
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(address_id) FROM address WHERE address_id = ?", addressId).Scan(&check)
	if err != nil {
		log.Printf("%s", err)
		return false, err
	}

	if check != 1 {
		return false, errors.New("Somethings wrong!")
	}

	return true, nil
}

func (a *Address) FindAddressById(addressId int) (*Address, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return &Address{}, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT address_id, province_id, city_id, sub_district_id, address, created_at, updated_at FROM address WHERE address_id = ?", addressId).
		Scan(&a.AddressId, &a.ProvinceId, &a.CityId, &a.SubDistrictId, &a.Address, &a.Audit.CreatedAt, &a.Audit.UpdatedAt)
	if err != nil {
		log.Printf("%s", err)
		return &Address{}, errors.New("Somethings wrong!")
	}

	if a == (&Address{}) {
		return &Address{}, errors.New(fmt.Sprintf("Can't find address with id: %d", addressId))
	}

	return a, nil
}

func (a *Address) SaveAddress() (*Address, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return &Address{}, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO address (province_id, city_id, sub_district_id, address, created_at) VALUES (?, ?, ?, ?, ?)",
		a.ProvinceId,
		a.CityId,
		a.SubDistrictId,
		a.Address,
		a.Audit.CreatedAt)
	if err != nil {
		log.Printf("%s", err)
		return &Address{}, errors.New("Somethings wrong!")
	}

	currentId, err := result.LastInsertId()
	if err != nil {
		log.Printf("%s", err)
		return &Address{}, errors.New("Somethings wrong!")
	}

	a.AddressId = int(currentId)
	return a, nil
}

func (a *Address) UpdateAddressById(addressId int) (*Address, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &Address{}, err
	}
	defer db.Close()

	result, err := db.Exec("UPDATE address SET address_id = ?, province_id = ?, city_id = ?, sub_district_id = ?, address = ?, created_at = ?, updated_at = ? WHERE address_id = ?",
		a.AddressId,
		a.ProvinceId,
		a.CityId,
		a.SubDistrictId,
		a.Address,
		a.Audit.CreatedAt,
		a.Audit.UpdatedAt,
		addressId)
	if err != nil {
		return &Address{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &Address{}, err
	}

	if rowsAffected != 1 {
		return &Address{}, errors.New("Somethings wrong!")
	}

	return a, nil
}

func (a *Address) DeleteAddressById(addressId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM address WHERE address_id = ?", addressId)
	if err != nil {
		return false, err
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

func (a *Address) FindAllAddress(limit int, offset int) ([]*Address, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return []*Address{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT address_id, province_id, city_id, sub_district_id, address, created_at, updated_at FROM address LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		log.Printf("%s", err)
		return []*Address{}, err
	}
	defer rows.Close()

	result := []*Address{}
	for rows.Next() {
		each := &Address{}
		err = rows.Scan(&each.AddressId, &each.ProvinceId, &each.CityId, &each.SubDistrictId, &each.Address, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*Address{}, err
		}

		result = append(result, each)
	}

	return result, nil
}
