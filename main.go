package main

import (
	"fmt"
	"log"

	"sales-backend/model"
	"github.com/joho/godotenv"
)

func init() {
	load := godotenv.Load()
	if load != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	// init1 := model.CategoryProduct{Category: "Health"}
	// init2 := model.CategoryProduct{Category: "Technology"}
	// init3 := model.CategoryProduct{Category: "Lifestyle"}

	// init1.SaveCategoryProduct()
	// init2.SaveCategoryProduct()
	// init3.SaveCategoryProduct()

	// fmt.Println(init1)
	// fmt.Println(init2)
	// fmt.Println(init3)

	repo := model.CategoryProduct{}
	listCategoryProduct, err := repo.FindAllCategoryProduct()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(listCategoryProduct); i++ {
		// fmt.Println(*listCategoryProduct[i])
		instance := *listCategoryProduct[i]
		instance.Category = "New Category Again Again"
		instance.UpdateCategoryProduct()
	}
}
