package model


type TransactionDetail struct {
	TransactionDetailId	string 	`json:"transactionDetailId"`
	TransactionId       	string 	`json:"transactionId"`
	ProductId           	string 	`json:"productId"`
	Price               	int64  	`json:"price"`
	Quantity            	int    	`json:"quantity"`
	Audit               	Audit 	`json:"audit"`
}
