package model

type Province struct {
	ProvinceId	int	`json:"provinceId"`
	Province   	string	`json:"province"`
	Audit		Audit	`json:"audit"`
}
