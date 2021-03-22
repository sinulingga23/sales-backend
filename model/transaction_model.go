package model

type Transaction struct {
	TransactionId	string 	`json:"transactionId"`
	CustomerId	string 	`json:"customerId"`
	EmployeeId	string 	`json:"employeeId"`
	Date		string 	`json:"date"`
	Audit		Audit 	`json:"audit"`
}
