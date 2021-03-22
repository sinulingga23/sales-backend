package model


type City struct {
	CityId		int	`json:"cityId"`
	ProvinceId	int	`json:"provinceId"`
	City		string	`json:"city"`
	Audit		Audit	`json:"audit"`
}
