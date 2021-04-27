package response

import (
	"sales-backend/model"
)

type ResponseGeneric struct {
	StatusCode 	int	`json:"statusCode"`
	Message 	string 	`json:"message"`
}

type ResponseErrors struct {
	StatusCode	int 	`json:"statusCode"`
	Message 	string 	`json:"message"`
	Errors		string 	`json:"erros"`
}

type ResponseProvince struct {
	StatusCode	int		`json:"statusCode"`
	Message 	string		`json:"message"`
	Province	model.Province	`json:"province"`
}

type ResponseProvinces struct {
	StatusCode	int 			`json:"statusCode"`
	Message		string			`json:"message"`
	Provinces 	[]*model.Province	`json:"provinces"`
}

type ResponseCategoryProduct struct {
	StatusCode 		int 			`json:"statusCode"`
	Message			string 			`json:"message"`
	CategoryProduct 	model.CategoryProduct	`json:"categoryProduct"`
}

type ResponseCategoryProducts struct {
	StatusCode 		int 				`json:"statusCode"`
	Message 		string				`json:"message"`
	CategoryProducts 	[]*model.CategoryProduct	`json:"categoryProducts"`
}

type ResponseCity struct {
	StatusCode	int 		`json:"statusCode"`
	Message		string		`json:"message"`
	City		model.City	`json:"city"`
}

type ResponseCities struct {
	StatusCode	int		`json:"statusCode"`
	Message		string 		`json:"message"`
	Cities		[]*model.City	`json:"cities"`
}

type ResponseProducts struct {
	StatusCode	int			`json:"statusCode"`
	Message		string			`json:"message"`
	Products	[]*model.Product	`json:"products"`
}
