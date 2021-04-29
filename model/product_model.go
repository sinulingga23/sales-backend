package model

import (
	"sales-backend/utility"
)

type Product struct {
	ProductId         string  `json:"productId"`
	CategoryProductId string  `json:"categoryProductId"`
	Name              string  `json:"product"`
	Unit              string  `json:"unit"`
	Price             float64 `json:"price"`
	Stock             int     `json:"stock"`
	Audit             Audit   `json:"audit"`
}

func (p *Product) GetNumberRecords() (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(product_id) FROM product").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}

func (p *Product) GetNumberRecordsByCategoryProductId(categoryProductId string) (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT COUNT(product_id) FROM product WHERE category_product_id = ?", categoryProductId).Scan(&numberRecords)
	if err != nil {
		return 0, err
	}

	return numberRecords, nil
}
