package model

type Audit struct {
	CreatedAt 				string		`json:"createdAt"`
	UpdatedAt 				string		`json:"updatedAt"`
}

type CategoryProduct struct {
	CategoryProductId		string		`json:"categoryProductId"`
	Category				string		`json:"category"`
	Audit					string		`json:"audit"`

}

type Product struct {
	ProductId				string		`json:"productId"`
	CategoryProductId		string		`json:"categoryProductId"`
	Name					string 		`json:"product"`
	Unit					int 		`json:"unit"`
	Price					int64		`json:"price"`
	Stock					int			`json:"stock"`
	Audit					string		`json:"audit"`
}

type Customer struct {
	CustomerId				string		`json:"customerId"`
	FirstName				string		`json:"firstName"`
	LastName				string		`json:"lastName"`
	Gender					string		`json:"gender"`
	Address					string		`json:"address"`
	City					string		`json:"city"`
	PhoneNumber				string		`json:"phoneNumber"`
	Audit					string		`json:"audit"`
}

type Employee struct {
	EmployeeId				string		`json:"employeeId"`
	FirstName				string		`json:"firstName"`
	LastName				string		`json:"lastName"`
	Address					string		`json:"address"`
	PhoneNumber				string		`json:"phoneNumber"`
	Audit					string		`json:"audit"`
}

type Transaction struct {
	TransactionId			string		`json:"transactionId"`
	CustomerId				string		`json:"customerId"`
	EmployeeId				string		`json:"employeeId"`
	Date					string		`json:"date"`
	Audit					string		`json:"audit"`
}

type TransacitionDetail struct {
	TransacitionDetailId	string		`json:"transactionDetailId"`
	TransactionId			string		`json:"transactionId"`
	ProductId				string		`json:"productId"`
	Price					int64		`json:"price"`
	Quantity				int			`json:"quantity"`
	Audit					string		`json:"audit"`
}
