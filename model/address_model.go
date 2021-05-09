package model

import (
	"fmt"
	"log"
	"errors"

	"sales-backend/utility"
)

type Address struct {
	AddressId	int	`json:"addressId"`
	ProvinceId	int 	`json:"provinceId"`
	CityId		int	`json:"cityId"`
	SubDistrictId	int	`json:"subDistrictId"`
	Address		string	`json:"address"`
	Audit		Audit	`json:"audit"`
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
