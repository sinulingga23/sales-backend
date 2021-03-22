package model

type Product struct {
	ProductId		string `json:"productId"`
	CategoryProductId 	string `json:"categoryProductId"`
	Name              	string `json:"product"`
	Unit              	string `json:"unit"`
	Price             	int64  `json:"price"`
	Stock             	int    `json:"stock"`
	Audit             	Audit  `json:"audit"`
}
