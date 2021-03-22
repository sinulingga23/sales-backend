package model


type Customer struct {
	CustomerId 	string 	`json:"customerId"`
	FirstName   	string 	`json:"firstName"`
	LastName    	string 	`json:"lastName"`
	Gender      	string 	`json:"gender"`
	Address     	Address `json:"address"`
	Email		string	`json:"email"`
	PhoneNumber 	string 	`json:"phoneNumber"`
	Audit      	Audit 	`json:"audit"`
}
