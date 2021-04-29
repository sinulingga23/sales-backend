package response

import (
	"sales-backend/model"
)

type ResponseGeneric struct {
	StatusCode 	int	`json:"statusCode"`
	Message 	string 	`json:"message"`
}

type InfoPagination struct {
	CurrentPage		int	`json:"currentPage"`
	RowsEachPage		int	`json:"rowsEachPage"`
	TotalPages		int	`json:"totalPages"`
}

type ResponseErrors struct {
	StatusCode	int 	`json:"statusCode"`
	Message 	string 	`json:"message"`
	Errors		string 	`json:"errors"`
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
	InfoPagination	InfoPagination		`json:"infoPagination"`
	NextPage	string 			`json:"nextPage"`
	PrevPage	string 			`json:"prevPage"`
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
	InfoPagination		InfoPagination			`json:"infoPagination"`
	NextPage		string				`json:"nextPage"`
	PrevPage		string				`json:"prevPage"`
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

type ResponseProductsByCategoryProductId struct {
	StatusCode		int			`json:"statusCode"`
	Message			string			`json:"message"`
	CategoryProductId	string			`json:"categoryProductId"`
	Products		[]*model.Product	`json:"products"`
	InfoPagination		InfoPagination		`json:"infoPagination"`
	NextPage		string			`json:"nextPage"`
	PrevPage		string			`json:"prevPage"`
}
