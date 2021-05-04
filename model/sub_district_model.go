package model

import (
	"sales-backend/utility"
)

type SubDistrict struct {
	SubDistrictId	int	`json:"subDistrictId"`
	CityId		int	`json:"cityId"`
	SubDistrict	string	`json:"subDistrict"`
	Audit		Audit	`json:"audit"`
}

func (sD *SubDistrict) GetNumberRecordsByCityId(cityId int) (int, error) {
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
