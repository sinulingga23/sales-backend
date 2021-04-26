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

type ResponseCategory struct {
	StatusCode 		int 			`json:"statusCode"`
	Message			string 			`json:"message"`
	CategoryProduct 	model.CategoryProduct	`json:"categoryProduct"`
}

type ResponseCategories struct {
	StatusCode 		int 				`json:"statusCode"`
	Message 		string				`json:"message"`
	CategoryProducts 	[]*model.CategoryProduct	`json:"categoryProducts"`
}
