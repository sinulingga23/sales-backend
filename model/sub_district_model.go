package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sinulingga23/sales-backend/utility"
)

type SubDistrict struct {
	SubDistrictId string `json:"subDistrictId"`
	CityId        string `json:"cityId"`
	SubDistrict   string `json:"subDistrict"`
	Audit         Audit  `json:"audit"`
}

func (sD *SubDistrict) GetNumberRecords() (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(sub_district_id) FROM sub_district").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}

func (sD *SubDistrict) GetNumberRecordsByCityId(cityId string) (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(sub_district_id) FROM sub_district WHERE city_id = ?", cityId).Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}

func (sD *SubDistrict) IsSubDistrictExistsById(subDistrictId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(sub_district_id) FROM sub_district WHERE sub_district_id = ?", subDistrictId).Scan(&check)
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}

	if check != 1 {
		return false, nil
	}

	return true, nil
}

func (sD *SubDistrict) SaveSubDistrict() (*SubDistrict, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	subDistrictId := uuid.NewString()
	_, err = db.Exec("INSERT INTO sub_district (sub_district_id, city_id, sub_district, created_at) VALUES (?, ?, ?)",
		subDistrictId,
		sD.CityId,
		sD.SubDistrict,
		sD.Audit.CreatedAt)
	if err != nil {
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}

	sD.SubDistrictId = subDistrictId
	return sD, nil
}

func (sD *SubDistrict) FindSubDistrictById(subDistrictId string) (*SubDistrict, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	err = db.QueryRow("SELECT sub_district_id, city_id, sub_district, created_at, updated_at FROM sub_district WHERE sub_district_id = ?", subDistrictId).
		Scan(&sD.SubDistrictId, &sD.CityId, &sD.SubDistrict, &sD.Audit.CreatedAt, &sD.Audit.UpdatedAt)
	if err != nil {
		log.Printf("%s", err)
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}

	if sD == (&SubDistrict{}) {
		return &SubDistrict{}, errors.New(fmt.Sprintf("Can't find sub-district with id: %v", subDistrictId))
	}

	return sD, nil
}

func (sD *SubDistrict) UpdateSubDistrictById(subDistrictId string) (*SubDistrict, error) {
	db, err := utility.ConnectDB()
	if err != nil || sD.SubDistrictId != subDistrictId {
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE sub_district SET sub_district_id = ?, city_id = ?, sub_district = ?, created_at = ?, updated_at = ? WHERE sub_district_id = ?",
		sD.SubDistrictId,
		sD.CityId,
		sD.SubDistrict,
		sD.Audit.CreatedAt,
		sD.Audit.UpdatedAt,
		subDistrictId)
	if err != nil {
		log.Printf("%s", err)
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s", err)
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		log.Printf("%s", err)
		return &SubDistrict{}, errors.New("Somethings wrong!")
	}

	return sD, nil
}

func (sD *SubDistrict) DeleteSubDistrictById(subDistrictId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM sub_district WHERE sub_district_id = ?", subDistrictId)
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

func (sD *SubDistrict) FindAllSubDistrict(limit int, offset int) ([]*SubDistrict, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*SubDistrict{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT sub_district_id, city_id, sub_district, created_at, updated_at FROM sub_district LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return []*SubDistrict{}, err
	}
	defer rows.Close()

	result := []*SubDistrict{}
	for rows.Next() {
		each := &SubDistrict{}
		err = rows.Scan(&each.SubDistrictId, &each.CityId, &each.SubDistrict, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*SubDistrict{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*SubDistrict{}, errors.New("Somethings wrong!")
	}

	return result, nil
}
