package model

type SubDistrict struct {
	SubDistrictId	int	`json:"subDistrictId"`
	CityId		int	`json:"cityId"`
	SubDistrict	string	`json:"subDistrict"`
	Audit		Audit	`json:"audit"`
}
