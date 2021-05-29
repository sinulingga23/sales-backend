package model

import (
	"log"
	"fmt"
	"errors"

	"sales-backend/utility"
)

type Product struct {
	ProductId		string	`json:"productId"`
	CategoryProductId	string  `json:"categoryProductId"`
	Name             	string  `json:"name"`
	Unit             	string  `json:"unit"`
	Price            	float64 `json:"price"`
	Stock            	int     `json:"stock"`
	AddStock		int 	`json:"addStock"` // Use this field for adding new stock
	Audit            	Audit   `json:"audit"`
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

func (p *Product) IsProductExistsById(productId string)  (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(product_id) FROM product WHERE product_id = ?", productId).Scan(&check)
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}

	if check != 1 {
		return false, nil
	}

	return true, err
}

func (p *Product) SaveProduct() (*Product, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	number := 0
	err = db.QueryRow("SELECT COUNT(product_id) FROM product").Scan(&number)
	if err != nil {
		log.Printf("%s", err)
		return &Product{}, errors.New("Something wrong!")
	}

	count := utility.DigitsCount(number)
	formatProductId := "PRD00000000"
	productId := "PRD"
	for i := 0; i < len(formatProductId)-count-5; i++ {
		productId += "0"
	}
	number += 1
	productId += fmt.Sprintf("%d", number)
	p.ProductId = productId

	_, err = db.Exec("INSERT INTO product (product_id, category_product_id, name, unit, price, stock, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.ProductId,
		p.CategoryProductId,
		p.Name,
		p.Unit,
		p.Price,
		p.Stock,
		p.Audit.CreatedAt)
	if err != nil {
		log.Printf("%s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}

	return p, nil
}

func (p *Product) FindProductById(productId string) (*Product, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	err = db.QueryRow("SELECT product_id, category_product_id, name, unit, price, stock, created_at, updated_at FROM product WHERE product_id = ?", productId).
		Scan(&p.ProductId, &p.CategoryProductId, &p.Name, &p.Unit, &p.Price, &p.Stock, &p.Audit.CreatedAt, &p.Audit.UpdatedAt)
	if err != nil {
		log.Printf("%s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}

	if p == (&Product{}) {
		return &Product{}, errors.New(fmt.Sprintf("Can't find Product with id: %d", productId))
	}

	return p, nil
}

func (p *Product) UpdateProductById(productId string) (*Product, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE product SET product_id = ?, category_product_id = ?, name = ?, unit = ?, price = ?, stock = ?, created_at = ?, updated_at = ? WHERE product_id = ?",
		p.ProductId,
		p.CategoryProductId,
		p.Name,
		p.Unit,
		p.Price,
		p.Stock,
		p.Audit.CreatedAt,
		p.Audit.UpdatedAt,
		productId)
	if err != nil {
		log.Printf("189: %s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("195: %s", err)
		return &Product{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		log.Printf("200 Err")
		return &Product{}, errors.New("Somethings wrong!")
	}

	return p, nil
}

func (p *Product) DeleteProductById(productId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM product WHERE product_id = ?", productId)
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s", err)
		return false, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return false, nil
	}

	return true, nil
}

func (p *Product) FindAllProduct(limit int, offset int) ([]*Product, error){
	db, err := utility.ConnectDB()
	if err != nil {
		log.Printf("%s", err)
		return []*Product{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT product_id, category_product_id, name, unit, price, stock, created_at, updated_at FROM product LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		log.Printf("%s", err)
		return []*Product{}, err
	}
	defer rows.Close()

	result := []*Product{}
	for rows.Next() {
		each := &Product{}
		err = rows.Scan(&each.ProductId, &each.CategoryProductId, &each.Name, &each.Unit, &each.Price, &each.Stock, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			return []*Product{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*Product{}, errors.New("Somethings wrong!")
	}

	return result, nil
}
