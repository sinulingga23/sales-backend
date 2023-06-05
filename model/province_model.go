package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	Province struct {
		ProvinceId string `json:"provinceId"`
		Province   string `json:"province"`
		Audit      Audit  `json:"audit"`
	}

	ProvinceRequest struct {
		Province string `json:"province"`
	}
)

type ProvinceRepository struct {
	db *sql.DB
}

func NewprovinceRepository(db *sql.DB) *ProvinceRepository {
	return &ProvinceRepository{db: db}
}

func (repository *ProvinceRepository) GetNumberRecords() (int, error) {

	numberRecords := 0
	err := repository.db.QueryRow("SELECT COUNT(province_id) FROM province").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}

func (repository *ProvinceRepository) IsProvinceExistsById(provinceId string) (bool, error) {
	check := 0
	err := repository.db.QueryRow("SELECT COUNT(province_id) FROM province WHERE province_id = ?", provinceId).Scan(&check)
	if err != nil {
		return false, err
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (repository *ProvinceRepository) SaveProvince(p Province) (*Province, error) {
	p.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")

	provinceId := uuid.NewString()
	_, err := repository.db.Exec("INSERT INTO province (province_id, province, created_at) VALUES (?, ?)",
		provinceId,
		p.Province,
		p.Audit.CreatedAt)

	if err != nil {
		return &Province{}, err
	}

	p.ProvinceId = provinceId
	return &p, nil
}

func (repository *ProvinceRepository) FindProvinceById(provinceId string) (*Province, error) {
	p := &Province{}

	err := repository.db.QueryRow("SELECT province_id, province, created_at, updated_at FROM province WHERE province_id = ?", provinceId).
		Scan(&p.ProvinceId, &p.Province, &p.Audit.CreatedAt, &p.Audit.UpdatedAt)
	if err != nil {
		return &Province{}, err
	}

	if p == (&Province{}) {
		return &Province{}, errors.New("The province can't found!")
	}

	return p, nil
}

func (repository *ProvinceRepository) UpdateProvinceById(provinceId string) (*Province, error) {
	p := &Province{}
	result, err := repository.db.Exec("UPDATE province SET province = ?, created_at = ?, updated_at = ? WHERE province_id = ?",
		p.Province,
		p.Audit.CreatedAt,
		p.Audit.UpdatedAt,
		provinceId)
	if err != nil {
		return &Province{}, err
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

func (repository *ProvinceRepository) DeleteProvinceById(provinceId string) (bool, error) {
	result, err := repository.db.Exec("DELETE FROM province WHERE province_id = ?", provinceId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected != 1 {
		return false, nil
	}

	return true, nil
}

func (repository *ProvinceRepository) FindAllProvince(limit int, offset int) ([]*Province, error) {
	rows, err := repository.db.Query("SELECT province_id, province, created_at, updated_at FROM province LIMIT ? OFFSET ?", limit, offset)
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

func (repository *ProvinceRepository) FindAllCityByProvinceId(provinceId string, limit int, offset int) ([]*City, error) {
	rows, err := repository.db.Query("SELECT c.city_id, c.province_id, c.city, c.created_at, c.updated_at FROM city c INNER JOIN province p ON c.province_id = p.province_id HAVING c.province_id = ? LIMIT ? OFFSET ?", provinceId, limit, offset)
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
