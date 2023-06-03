package model

import (
	"errors"
	"fmt"
	"log"
	"sales-backend/utility"
)

type CategoryProduct struct {
	CategoryProductId string `json:"categoryProductId"`
	Category          string `json:"category"`
	Audit             Audit  `json:"audit"`
}

func (c *CategoryProduct) IsCategoryProductExistsById(categoryProductId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	check := 0
	err = db.QueryRow("SELECT COUNT(category_product_id) FROM category_product WHERE category_product_id = ?", categoryProductId).Scan(&check)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if check != 1 {
		return false, nil
	}
	return true, nil
}

func (c *CategoryProduct) SaveCategoryProduct() (*CategoryProduct, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	// Get the number rows of category_product
	number := 0
	err = db.QueryRow("SELECT count(category_product_id) FROM category_product").Scan(&number)
	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}

	// generate the categeory_product_id
	count := utility.DigitsCount(number)
	formatCategoryProductId := "CTG00000000"
	categoryProductId := "CTG"
	for i := 0; i < len(formatCategoryProductId)-count-5; i++ {
		categoryProductId += "0"
	}
	number += 1
	categoryProductId += fmt.Sprintf("%d", number)
	c.CategoryProductId = categoryProductId

	// Query add category
	_, err = db.Exec("INSERT INTO category_product (category_product_id, category, created_at) VALUES (?, ?, ?)",
		c.CategoryProductId,
		c.Category,
		c.Audit.CreatedAt)

	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}
	return c, nil
}

func (c *CategoryProduct) FindCategoryProductById(categoryProductId string) (*CategoryProduct, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &CategoryProduct{}, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT category_product_id, category, created_at, updated_at FROM category_product WHERE category_product_id = ?", categoryProductId).Scan(&c.CategoryProductId, &c.Category, &c.Audit.CreatedAt, &c.Audit.UpdatedAt)
	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}

	if c == (&CategoryProduct{}) {
		return &CategoryProduct{}, errors.New(fmt.Sprintf("Can't find category product with id: %s", categoryProductId))
	}

	return c, nil
}

func (c *CategoryProduct) UpdateCategoryProduct() (*CategoryProduct, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("UPDATE category_product SET category_product_id = ?, category = ?, created_at = ?, updated_at = ? WHERE category_product_id = ?",
		c.CategoryProductId,
		c.Category,
		c.Audit.CreatedAt,
		c.Audit.UpdatedAt,
		c.CategoryProductId)
	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}

	if rowsAffected != 1 {
		return &CategoryProduct{}, errors.New("Somethings wrong!")
	}

	return c, nil
}

func (c *CategoryProduct) DeleteCategoryProductById(categoryProductId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM category_product WHERE category_product_id = ?", categoryProductId)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}
	if rowsAffected != 1 {
		return false, nil
	}
	return true, nil
}

func (c *CategoryProduct) GetNumberRecords() (int, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	numberRecords := 0
	err = db.QueryRow("SELECT count(category_product_id) FROM category_product").Scan(&numberRecords)
	if err != nil {
		return 0, err
	}
	return numberRecords, nil
}

func (c *CategoryProduct) FindAllCategoryProduct(limit int, offset int) ([]*CategoryProduct, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*CategoryProduct{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT category_product_id, category, created_at, updated_at FROM category_product LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return []*CategoryProduct{}, err
	}
	defer rows.Close()

	result := []*CategoryProduct{}
	for rows.Next() {
		each := &CategoryProduct{}
		err = rows.Scan(&each.CategoryProductId, &each.Category, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)

		if err != nil {
			return []*CategoryProduct{}, err
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*CategoryProduct{}, err
	}

	return result, nil
}

func (c *CategoryProduct) FindAllProductByCategoryProductId(categoryProductId string, limit int, offset int) ([]*Product, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*Product{}, errors.New("Somethings wrong!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT p.product_id, p.category_product_id, p.name, p.unit, p.price, p.stock, p.created_at, p.updated_at FROM product p INNER JOIN category_product cp ON p.category_product_id = cp.category_product_id HAVING p.category_product_id = ? LIMIT ? OFFSET ?", categoryProductId, limit, offset)
	if err != nil {
		log.Printf("query: %v", err)
		return []*Product{}, errors.New("Somethings wrong!")
	}
	defer rows.Close()

	result := []*Product{}
	for rows.Next() {
		each := &Product{}
		err = rows.Scan(&each.ProductId, &each.CategoryProductId, &each.Name, &each.Unit, &each.Price, &each.Stock, &each.Audit.CreatedAt, &each.Audit.UpdatedAt)
		if err != nil {
			log.Printf("parsing: %v", err)
			return []*Product{}, errors.New("Somethings wrong!")
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return []*Product{}, errors.New("Somethings wrong!")
	}
	return result, nil
}
