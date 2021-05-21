package model

import (
	"fmt"
	"log"
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
		return false, nil
	}

	return true, nil
}

func (c *City) SaveCity() (*City, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO city (province_id, city, created_at) VALUES (?, ?, ?)",
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
		c.Audit.UpdatedAt,
		cityId)
	if err != nil {
		log.Printf("%s", err)
		return &City{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s", err)
		return &City{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		log.Printf("%s", err)
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
		return false, nil
	}

	return true, nil
}

func (c *City) FindAllCity(limit int, offset int) ([]*City, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*City{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT city_id, province_id, city, created_at, updated_at FROM city LIMIT ? OFFSET ?", limit, offset)
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

func (c *City) FindAllSubDistrictByCityId(cityId int, limit int, offset int) ([]*SubDistrict, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%v", err)
		return []*SubDistrict{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT sb.sub_district_id, sb.city_id, sb.sub_district, sb.created_at, sb.updated_at FROM sub_district sb INNER JOIN city c ON sb.city_id = c.city_id HAVING sb.city_id = ? LIMIT ? OFFSET ?", cityId, limit, offset)
	if err != nil {
		log.Printf("%v", err)
		return []*SubDistrict{}, errors.New("Somethings wrong!")
	}
	defer rows.Close()

	result := []*SubDistrict{}
	for rows.Next() {
		each := &SubDistrict{}
		err = rows.Scan(&each.SubDistrictId, &each.CityId, &each.SubDistrict, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			log.Printf("%v", err)
			return []*SubDistrict{}, errors.New("Somethings wrong!")
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		log.Printf("%v", err)
		return []*SubDistrict{}, errors.New("Somethings wrong!")
	}

	return result, nil
}

func (c *City) GetNumberRecords() (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(city_id) FROM city").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
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
