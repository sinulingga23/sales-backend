package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type (
	City struct {
		CityId     string `json:"cityId"`
		ProvinceId string `json:"provinceId"`
		City       string `json:"city"`
		Audit      Audit  `json:"audit"`
	}

	CityRequest struct {
		ProvinceId string `json:"provinceId"`
		City       string `json:"city"`
	}
)

type CityRepository struct {
	db *sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (repository *CityRepository) IsCityExistsById(cityId string) (bool, error) {

	check := 0
	err := repository.db.QueryRow("SELECT COUNT(city_id) FROM city WHERE city_id = ?", cityId).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (repository *CityRepository) SaveCity(c City) (*City, error) {

	cityId := uuid.NewString()
	c.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	_, err := repository.db.Exec("INSERT INTO city (city_id, province_id, city, created_at) VALUES (?, ?, ?)",
		cityId,
		c.ProvinceId,
		c.City,
		c.Audit.CreatedAt)
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	c.CityId = cityId
	return &c, nil
}

func (repository *CityRepository) FindCityById(cityId string) (*City, error) {
	c := &City{}

	err := repository.db.QueryRow("SELECT city_id, province_id, city, created_at, updated_at FROM city WHERE city_id = ?", cityId).
		Scan(&c.CityId, &c.ProvinceId, &c.City, &c.Audit.CreatedAt, &c.Audit.UpdatedAt)
	if err != nil {
		return &City{}, errors.New("Somethings wrong!")
	}

	if c == (&City{}) {
		return &City{}, errors.New(fmt.Sprintf("Can't find city with id: %v", cityId))
	}

	return c, nil
}

func (repository *CityRepository) UpdateCityById(cityId string, c City) (*City, error) {
	tx, err := repository.db.Begin()
	if err != nil {
		return &City{}, err
	}

	result, err := tx.Exec("UPDATE city SET province_id = ?, city = ?, created_at = ?, updated_at = ? WHERE city_id = ?",
		c.ProvinceId,
		c.City,
		c.Audit.CreatedAt,
		c.Audit.UpdatedAt,
		cityId)
	if err != nil {
		tx.Rollback()
		return &City{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return &City{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		tx.Rollback()
		return &City{}, errors.New("Somethings wrong!")
	}

	updatedCity := &City{}
	row := tx.QueryRow("SELECT city_id, province_id, city, created_at, updated_at from city where city_id = ?", cityId)
	err = row.Scan(
		&updatedCity.CityId,
		&updatedCity.ProvinceId,
		&updatedCity.City,
		&updatedCity.Audit.CreatedAt,
		&updatedCity.Audit.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return &City{}, err
	}

	return updatedCity, nil
}

func (repository *CityRepository) DeleteCityById(cityId string) (bool, error) {
	result, err := repository.db.Exec("DELETE FROM city WHERE city_id = ?", cityId)
	if err != nil {
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

func (repository *CityRepository) FindAllCity(limit int, offset int) ([]*City, error) {
	rows, err := repository.db.Query("SELECT city_id, province_id, city, created_at, updated_at FROM city LIMIT ? OFFSET ?", limit, offset)
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

func (repository *CityRepository) FindAllSubDistrictByCityId(cityId string, limit int, offset int) ([]*SubDistrict, error) {
	rows, err := repository.db.Query("SELECT sb.sub_district_id, sb.city_id, sb.sub_district, sb.created_at, sb.updated_at FROM sub_district sb INNER JOIN city c ON sb.city_id = c.city_id HAVING sb.city_id = ? LIMIT ? OFFSET ?", cityId, limit, offset)
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

func (repository *CityRepository) GetNumberRecords() (int, error) {
	numberRecords := 0
	err := repository.db.QueryRow("SELECT COUNT(city_id) FROM city").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}

func (repository *CityRepository) GetNumberRecordsByProvinceId(provinceId string) (int, error) {
	numberRecords := 0
	err := repository.db.QueryRow("SELECT COUNT(city_id) FROM city WHERE province_id = ?", provinceId).Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}
