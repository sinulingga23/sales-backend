package model

type Audit struct {
	CreatedAt	string	`json:"createdAt"`
	UpdatedAt 	string 	`json:"updatedAt"`
}

type CategoryProduct struct {
	CategoryProductId	string `json:"categoryProductId"`
	Category          	string `json:"category"`
	Audit             	Audit  `json:"audit"`
}

type Product struct {
	ProductId		string `json:"productId"`
	CategoryProductId 	string `json:"categoryProductId"`
	Name              	string `json:"product"`
	Unit              	string `json:"unit"`
	Price             	int64  `json:"price"`
	Stock             	int    `json:"stock"`
	Audit             	Audit  `json:"audit"`
}

type Province struct {
	ProvinceId	int	`json:"provinceId"`
	Province   	string	`json:"province"`
	Audit		Audit	`json:"audit"`
}

type City struct {
	CityId		int	`json:"cityId"`
	ProvinceId	int	`json:"provinceId"`
	City		string	`json:"city"`
	Audit		Audit	`json:"audit"`
}

type SubDistrict struct {
	SubDistrictId	int	`json:"subDistrictId"`
	CityId		int	`json:"cityId"`
	SubDistrict	string	`json:"subDistrict"`
	Audit		Audit	`json:"audit"`
}

type Address struct {
	AddressId	int	`json:"addressId"`
	ProvinceId	int 	`json:"provinceId"`
	CityId		int	`json:"cityId"`
	SubDistrictId	int	`json:"subDistrictId"`
	Address		string	`json:"address"`
	Audit		Audit	`json:"audit"`
}

type Customer struct {
	CustomerId 	string 	`json:"customerId"`
	FirstName   	string 	`json:"firstName"`
	LastName    	string 	`json:"lastName"`
	Gender      	string 	`json:"gender"`
	Address     	Address `json:"address"`
	City        	string 	`json:"city"`
	PhoneNumber 	string 	`json:"phoneNumber"`
	Audit      	Audit 	`json:"audit"`
}

type Employee struct {
	EmployeeId	string	`json:"employeeId"`
	FirstName	string 	`json:"firstName"`
	LastName	string 	`json:"lastName"`
	Address		Address `json:"address"`
	PhoneNumber	string 	`json:"phoneNumber"`
	Audit       	Audit 	`json:"audit"`
}

type Transaction struct {
	TransactionId	string 	`json:"transactionId"`
	CustomerId	string 	`json:"customerId"`
	EmployeeId	string 	`json:"employeeId"`
	Date		string 	`json:"date"`
	Audit		Audit 	`json:"audit"`
}

type TransactionDetail struct {
	TransactionDetailId	string 	`json:"transactionDetailId"`
	TransactionId       	string 	`json:"transactionId"`
	ProductId           	string 	`json:"productId"`
	Price               	int64  	`json:"price"`
	Quantity            	int    	`json:"quantity"`
	Audit               	Audit 	`json:"audit"`
}
