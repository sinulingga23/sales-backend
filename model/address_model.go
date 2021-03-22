package model


type Address struct {
	AddressId	int	`json:"addressId"`
	ProvinceId	int 	`json:"provinceId"`
	CityId		int	`json:"cityId"`
	SubDistrictId	int	`json:"subDistrictId"`
	Address		string	`json:"address"`
	Audit		Audit	`json:"audit"`
}
