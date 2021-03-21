package main

import (
	"fmt"

	"sales-backend/model"
)


func main() {
	opt := model.CategoryProduct{}
	instance, err := opt.FindCategoryProductById("CTG0000001")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*instance.Audit.UpdatedAt)
}
