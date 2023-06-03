package modeltest

import (
	"fmt"
	"log"
	"testing"

	"github.com/sinulingga23/sales-backend/model"

	"github.com/stretchr/testify/assert"
)

func TestIsCategoryProductExistsById_Exists(t *testing.T) {
	assert := assert.New(t)

	categoryProductModel := model.CategoryProduct{}

	expectedResult1 := true
	expectedResult2 := true
	expectedResult3 := true

	actualResult1, err := categoryProductModel.IsCategoryProductExistsById("CTG0000001")
	if err != nil {
		log.Printf("%s", err)
	}

	actualResult2, err := categoryProductModel.IsCategoryProductExistsById("CTG0000002")
	if err != nil {
		log.Printf("%s", err)
	}

	actualResult3, err := categoryProductModel.IsCategoryProductExistsById("CTG0000003")
	if err != nil {
		log.Printf("%s", err)
	}

	assert.Equal(expectedResult1, actualResult1, "They should be equal")
	assert.Equal(expectedResult2, actualResult2, "They should be equal")
	assert.Equal(expectedResult3, actualResult3, "They should be equal")
}

func TestIsCategoryProductExistsByIdById_NotExists(t *testing.T) {
	assert := assert.New(t)

	categoryProductModel := model.CategoryProduct{}

	categoryProductId1 := "CTG0000010"
	categoryProductId2 := "CTG0000011"
	categoryProductId3 := "CTG0000012"

	expectedMessageResult1 := fmt.Sprintf("Can't find category product with id: %s", categoryProductId1)
	expectedMessageResult2 := fmt.Sprintf("Can't find category product with id: %s", categoryProductId2)
	expectedMessageResult3 := fmt.Sprintf("Can't find category product with id: %s", categoryProductId3)

	actualMessageResult1 := ""
	actualMessageResult2 := ""
	actualMessageResult3 := ""

	actualValueResult1, err1 := categoryProductModel.IsCategoryProductExistsById("CTG0000010")
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualValueResult2, err2 := categoryProductModel.IsCategoryProductExistsById("CTG0000011")
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualValueResult3, err3 := categoryProductModel.IsCategoryProductExistsById("CTG0000012")
	if err3 != nil {
		actualMessageResult3 = fmt.Sprintf("%s", err3)
	}

	assert.Equal(expectedMessageResult1, actualMessageResult1, "They should be equal")
	assert.Equal(expectedMessageResult2, actualMessageResult2, "They should be equal")
	assert.Equal(expectedMessageResult3, actualMessageResult3, "They should be equal")

	assert.NotEqual(true, actualValueResult1, "They should not be equal")
	assert.NotEqual(true, actualValueResult2, "They should not be equal")
	assert.NotEqual(true, actualValueResult3, "They should not be equal")
}

func TestIsCategoryProductExistsById_InputEmpty(t *testing.T) {
	assert := assert.New(t)

	categoryProductModel := model.CategoryProduct{}

	categoryProductId1 := ""
	categoryProductId2 := ""
	categoryProductId3 := ""

	expectedMessageResult1 := "CategoryProductId can't be empty"
	expectedMessageResult2 := "CategoryProductId can't be empty"
	expectedMessageResult3 := "CategoryProductId can't be empty"

	actualMessageResult1 := ""
	actualMessageResult2 := ""
	actualMessageResult3 := ""

	actualValueResult1, err1 := categoryProductModel.IsCategoryProductExistsById(categoryProductId1)
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualValueResult2, err2 := categoryProductModel.IsCategoryProductExistsById(categoryProductId2)
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualValueResult3, err3 := categoryProductModel.IsCategoryProductExistsById(categoryProductId3)
	if err3 != nil {
		actualMessageResult3 = fmt.Sprintf("%s", err3)
	}

	assert.Equal(expectedMessageResult1, actualMessageResult1, "They should be equal")
	assert.Equal(expectedMessageResult2, actualMessageResult2, "They should be equal")
	assert.Equal(expectedMessageResult3, actualMessageResult3, "They should be equal")

	assert.NotEqual(true, actualValueResult1, "They should not be equal")
	assert.NotEqual(true, actualValueResult2, "They should not be equal")
	assert.NotEqual(true, actualValueResult3, "They should not be equal")
}

func TestSaveCategoryProduct_Empty(t *testing.T) {
	assert := assert.New(t)

	expectedMessageResult1 := "Category name can't be empty"
	expectedMessageResult2 := "Category name can't be empty"
	expectedMessageResult3 := "Category name can't be empty"

	actualMessageResult1 := ""
	actualMessageResult2 := ""
	actualMessageResult3 := ""

	saveModel1 := model.CategoryProduct{}
	saveModel1.Category = ""
	actualModel1, err1 := saveModel1.SaveCategoryProduct()
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	saveModel2 := model.CategoryProduct{}
	actualModel2, err2 := saveModel2.SaveCategoryProduct()
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	saveModel3 := model.CategoryProduct{}
	saveModel3.Category = "    "
	actualModel3, err3 := saveModel3.SaveCategoryProduct()
	if err3 != nil {
		actualMessageResult3 = fmt.Sprintf("%s", err3)
	}

	assert.Equal(expectedMessageResult1, actualMessageResult1, "They should be equal")
	assert.Equal(expectedMessageResult2, actualMessageResult2, "They should be equal")
	assert.Equal(expectedMessageResult3, actualMessageResult3, "They should be equal")

	assert.Equal(&model.CategoryProduct{}, actualModel1, "They should be equal")
	assert.Equal(&model.CategoryProduct{}, actualModel2, "They should be equal")
	assert.Equal(&model.CategoryProduct{}, actualModel3, "They should be equal")
}

func TestSaveCategoryProduct_Success(t *testing.T) {
	assert := assert.New(t)

	saveModel1 := model.CategoryProduct{}
	saveModel1.Category = "Technology"
	actualModel1, err1 := saveModel1.SaveCategoryProduct()
	if err1 != nil {
		log.Printf("%s", err1)
	}

	saveModel2 := model.CategoryProduct{}
	saveModel2.Category = "Health"
	actualModel2, err2 := saveModel2.SaveCategoryProduct()
	if err2 != nil {
		log.Printf("%s", err2)
	}

	saveModel3 := model.CategoryProduct{}
	saveModel3.Category = "Sport"
	actualModel3, err3 := saveModel3.SaveCategoryProduct()
	if err3 != nil {
		log.Printf("%s", err3)
	}

	assert.NotEqual(&model.CategoryProduct{}, actualModel1, "They should not be equal")
	assert.NotEqual(&model.CategoryProduct{}, actualModel2, "They should not be equal")
	assert.NotEqual(&model.CategoryProduct{}, actualModel3, "They should not be equal")
}

func TestFindCategoryProductById_Empty(t *testing.T) {
	assert := assert.New(t)

	categoryProductModel := model.CategoryProduct{}

	expectedMessageResult1 := "CategoryProductId can't be empty"
	expectedMessageResult2 := "CategoryProductId can't be empty"
	expectedMessageResult3 := "CategoryProductId can't be empty"

	actualMessageResult1 := ""
	actualMessageResult2 := ""
	actualMessageResult3 := ""

	actualModel1, err1 := categoryProductModel.FindCategoryProductById("   ")
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualModel2, err2 := categoryProductModel.FindCategoryProductById("")
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualModel3, err3 := categoryProductModel.FindCategoryProductById("        ")
	if err3 != nil {
		actualMessageResult3 = fmt.Sprintf("%s", err3)
	}

	assert.Equal(expectedMessageResult1, actualMessageResult1, "They should be equal")
	assert.Equal(expectedMessageResult2, actualMessageResult2, "They should be equal")
	assert.Equal(expectedMessageResult3, actualMessageResult3, "They should be equal")

	assert.Equal(&model.CategoryProduct{}, actualModel1, "They should be equal")
	assert.Equal(&model.CategoryProduct{}, actualModel2, "They should be equal")
	assert.Equal(&model.CategoryProduct{}, actualModel3, "They should be equal")
}

func TestFindCategoryProductById_NotFound(t *testing.T) {
	assert := assert.New(t)

	categoryProductModel := model.CategoryProduct{}

	categoryProductId1 := "CTG00000101"
	categoryProductId2 := "CTG00000102"
	categoryProductId3 := "CTG00000103"

	expectedMessageResult1 := fmt.Sprintf("Can't find category product with id: %s", categoryProductId1)
	expectedMessageResult2 := fmt.Sprintf("Can't find category product with id: %s", categoryProductId2)
	expectedMessageResult3 := fmt.Sprintf("Can't find category product with id: %s", categoryProductId3)

	actualMessageResult1 := ""
	actualMessageResult2 := ""
	actualMessageResult3 := ""

	actualModel1, err1 := categoryProductModel.FindCategoryProductById(categoryProductId1)
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualModel2, err2 := categoryProductModel.FindCategoryProductById(categoryProductId2)
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualModel3, err3 := categoryProductModel.FindCategoryProductById(categoryProductId3)
	if err3 != nil {
		actualMessageResult3 = fmt.Sprintf("%s", err3)
	}

	assert.Equal(expectedMessageResult1, actualMessageResult1, "They should be equal")
	assert.Equal(expectedMessageResult2, actualMessageResult2, "They should be equal")
	assert.Equal(expectedMessageResult3, actualMessageResult3, "They should be equal")

	assert.Equal(&model.CategoryProduct{}, actualModel1, "They should be equal")
	assert.Equal(&model.CategoryProduct{}, actualModel2, "They should be equal")
	assert.Equal(&model.CategoryProduct{}, actualModel3, "They should be equal")
}

func TestFindCategoryProductById_Found(t *testing.T) {
	assert := assert.New(t)

	categoryProductModel1 := model.CategoryProduct{}
	categoryProductModel2 := model.CategoryProduct{}
	categoryProductModel3 := model.CategoryProduct{}

	categoryProductId1 := "CTG00000017"
	categoryProductId2 := "CTG0000003"
	categoryProductId3 := "CTG00000015"

	expectedMessageResult1 := ""
	expectedMessageResult2 := ""
	expectedMessageResult3 := ""

	actualMessageResult1 := ""
	actualMessageResult2 := ""
	actualMessageResult3 := ""

	expectedModel1 := &model.CategoryProduct{
		CategoryProductId: categoryProductId1,
		Category:          "Technology",
		Audit: model.Audit{
			CreatedAt: "2021-03-26 12:19:12",
			UpdatedAt: nil,
		},
	}

	updatedAt2 := "2021-03-22 13:42:01"
	expectedModel2 := &model.CategoryProduct{
		CategoryProductId: categoryProductId2,
		Category:          "New Category Again Again",
		Audit: model.Audit{
			CreatedAt: "2021-03-22 12:59:12",
			UpdatedAt: &updatedAt2,
		},
	}

	expectedModel3 := &model.CategoryProduct{
		CategoryProductId: categoryProductId3,
		Category:          "Health",
		Audit: model.Audit{
			CreatedAt: "2021-03-26 12:15:12",
			UpdatedAt: nil,
		},
	}

	actualModel1, err1 := categoryProductModel1.FindCategoryProductById(categoryProductId1)
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualModel2, err2 := categoryProductModel2.FindCategoryProductById(categoryProductId2)
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualModel3, err3 := categoryProductModel3.FindCategoryProductById(categoryProductId3)
	if err2 != nil {
		actualMessageResult3 = fmt.Sprintf("%s", err3)
	}

	assert.Equal(expectedMessageResult1, actualMessageResult1, "They should be equal")
	assert.Equal(expectedMessageResult2, actualMessageResult2, "They should be equal")
	assert.Equal(expectedMessageResult3, actualMessageResult3, "They should be equal")

	assert.Equal(expectedModel1, actualModel1, "They should be equal")
	assert.Equal(expectedModel2, actualModel2, "They should be equal")
	assert.Equal(expectedModel3, actualModel3, "They should be equal")
}
