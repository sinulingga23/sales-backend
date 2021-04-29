package model

import (
	"fmt"
	"errors"

	"sales-backend/utility"
)


type City struct {
	CityId		int	`json:"cityId"`
	ProvinceId	int	`json:"provinceId"`
	City		string	`json:"city"`
	Audit		Audit	`json:"audit"`
}

func (c *City) IsCityExistsById(cityId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0;
	err = db.QueryRow("SELECT COUNT(city_id) FROM city WHERE city_id = ?", cityId).Scan(&check)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if check != 1 {
		return false, errors.New(fmt.Sprintf("City with id %d is not exists.", cityId))
	}

	return true, nil
}

func (c *City) SaveCity() (*City, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO (province_id, city, created_at) VALUES (?, ?, ?)",
		c.ProvinceId,
		c.City,
		c.Audit.CreatedAt)
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	currentId, err := result.LastInsertId()
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	c.CityId = int(currentId)
	return c, nil
}

func (c *City) FindCityById(cityId int) (*City, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	err = db.QueryRow("SELECT city_id, province_id, city, created_at, updated_at FROM city WHERE city_id = ?", cityId).
		Scan(&c.CityId, &c.ProvinceId, &c.City, &c.Audit.CreatedAt, &c.Audit.UpdatedAt)
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	if c == (&City{}) {
		return &City{}, errors.New(fmt.Sprintf("Can't find city with id: %d", cityId))
	}

	return c, nil
}

func (c *City) UpdateCityById(cityId int) (*City, error) {
	db, err := utility.ConnectDB()
	if err != nil || c.CityId != cityId {
		return &City{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE city SET province_id = ?, city = ?, created_at = ?, updated_at = ? WHERE city_id = ?",
		c.ProvinceId,
		c.City,
		c.Audit.CreatedAt,
		c.Audit.UpdatedAt)
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return &City{}, errors.New("Somethings wrong!")
	}

	return c, nil
}

func (c *City) DeleteCityById(cityId int) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM city WHERE city_id = ?", cityId)
	if err != nil{
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

func (c *City) FindAllCity() ([]*City, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*City{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT city_id, province_id, city, created_at, updated_at FROM city")
	if err != nil {
		return []*City{}, errors.New("Somethings wrong!")
	}
	defer rows.Close()

	result := []*City{}
	for rows.Next() {
		each := &City{}
		err = rows.Scan(&each.CityId, &each.ProvinceId, &each.City, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*City{}, errors.New("Somethings wrong!")
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*City{}, errors.New("Somethings wrong!")
	}

	return result, nil
}

func (c *City) GetNumberRecordsByProvinceId(provinceId int) (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(city_id) FROM city WHERE province_id = ?", provinceId).Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}
