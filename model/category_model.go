package model

import (
	"fmt"
	"errors"
	"time"
	"sales-backend/utility"
)


type CategoryProduct struct {
	CategoryProductId	string `json:"categoryProductId"`
	Category          	string `json:"category"`
	Audit             	Audit  `json:"audit"`
}

func (c *CategoryProduct) IsCategoryProductExistsById(categoryProductId string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	if categoryProductId == "" {
		return false, errors.New("CategoryProductId can't be empty")
	}

	check := 0
	err = db.QueryRow("SELECT COUNT(category_product_id) FROM category_product WHERE category_product_id = ?",
		categoryProductId).Scan(&check)
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if check != 0 {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("Can't find category product with id: %s", categoryProductId))
}

func (c *CategoryProduct) SaveCategoryProduct() (*CategoryProduct, error) {
	// Make sure the important field is not empty
	if c.Category == "" {
		return &CategoryProduct{}, errors.New("Category name can't be empty")
	}

	// If the CreatedAt field is empty, then set the field using the current time
	if c.Audit.CreatedAt == "" {
		c.Audit.CreatedAt = time.Now().Format("2006-01-02 15:05:03")
	}

	db, err := utility.ConnectDB()
	if err != nil {
		return &CategoryProduct{}, err
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
	for i := 0; i<len(formatCategoryProductId)-count-5; i++ {
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
	isThere , err := c.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		return &CategoryProduct{}, err
	}


	if isThere {
		db, err := utility.ConnectDB()
		if err != nil {
			return &CategoryProduct{}, err
		}
		defer db.Close()

		err = db.QueryRow("SELECT category_product_id, category, created_at, updated_at FROM category_product WHERE category_product_id = ? ",categoryProductId).
			Scan(&c.CategoryProductId,&c.Category, &c.Audit.CreatedAt, &c.Audit.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return &CategoryProduct{}, errors.New("Somethings wrong!")
		}
	}

	return c, nil
}

func (c *CategoryProduct) UpdateCategoryProduct() (*CategoryProduct, error) {
	if c.CategoryProductId == "" {
		return &CategoryProduct{}, errors.New("Id can't be empty")
	} else if c.Category == ""{
		return &CategoryProduct{}, errors.New("Category name can't be empty")
	} else if c.Audit.CreatedAt == "" {
		return &CategoryProduct{}, errors.New("CreatedAt can't be empty")
	}

	timestamp := time.Now().Format("2006-01-02 15:05:03")
	c.Audit.UpdatedAt = &timestamp

	db, err := utility.ConnectDB()
	if err != nil {
		return &CategoryProduct{}, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE category_product SET category_product_id = ?, category = ?, created_at = ?, updated_at = ? WHERE category_product_id = ?",
		c.CategoryProductId,
		c.Category,
		c.Audit.CreatedAt,
		c.Audit.UpdatedAt,
		c.CategoryProductId)

	if err != nil {
		return &CategoryProduct{}, err
	}

	return c, nil
}

func (c *CategoryProduct) DeleteCategoryProductById(categoryProductId string) (bool, error) {
	isThere, err := c.IsCategoryProductExistsById(categoryProductId)
	if err != nil {
		return false, err
	}

	if isThere {
		db, err := utility.ConnectDB()
		if err != nil {
			return false, err
		}
		defer db.Close()

		_, err = db.Exec("DELETE FROM category_product WHERE category_product_id = ?", categoryProductId)
		if err != nil {
			return false, errors.New("Somethings wrong!")
		}
	}

	return true, nil
}

func (c *CategoryProduct) FindAllCategoryProduct() ([]*CategoryProduct, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return []*CategoryProduct{}, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT category_product_id, category, created_at, updated_at FROM category_product")
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

	if err = rows.Err(); err != nil  {
		return []*CategoryProduct{}, err
	}

	return result, nil
}
