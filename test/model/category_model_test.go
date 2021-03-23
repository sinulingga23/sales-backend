package modeltest


import (
	"log"
	"fmt"
	"testing"

	"sales-backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
)

func loadEnv() {
	load := godotenv.Load()
	if load != nil {
		log.Fatal("Error loading .env file")
	}
}

// TODO: TestTestIsCategoryProductExists_InputNotString

func TestIsCategoryProductExists_Exists(t *testing.T) {
	loadEnv()

	assert := assert.New(t)

	categoryProductModel := model.CategoryProduct{}

	expectedResult1 := true
	expectedResult2 := true
	expectedResult3 := true

	actualResult1, err := categoryProductModel.IsCategoryProductExists("CTG0000001")
	if err != nil {
		log.Printf("%s", err)
	}

	actualResult2, err := categoryProductModel.IsCategoryProductExists("CTG0000002")
	if err != nil {
		log.Printf("%s", err)
	}

	actualResult3, err := categoryProductModel.IsCategoryProductExists("CTG0000003")
	if err != nil {
		log.Printf("%s", err)
	}

	assert.Equal(expectedResult1, actualResult1, "They should be equal")
	assert.Equal(expectedResult2, actualResult2, "They should be equal")
	assert.Equal(expectedResult3, actualResult3, "They should be equal")
}

func TestIsCategoryProductExists_NotExists(t *testing.T) {
	loadEnv()

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

	actualValueResult1, err1 := categoryProductModel.IsCategoryProductExists("CTG0000010")
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualValueResult2, err2 := categoryProductModel.IsCategoryProductExists("CTG0000011")
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualValueResult3, err3 := categoryProductModel.IsCategoryProductExists("CTG0000012")
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

func TestIsCategoryProductExists_InputEmpty(t *testing.T) {
	loadEnv()

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

	actualValueResult1, err1 := categoryProductModel.IsCategoryProductExists(categoryProductId1)
	if err1 != nil {
		actualMessageResult1 = fmt.Sprintf("%s", err1)
	}

	actualValueResult2, err2 := categoryProductModel.IsCategoryProductExists(categoryProductId2)
	if err2 != nil {
		actualMessageResult2 = fmt.Sprintf("%s", err2)
	}

	actualValueResult3, err3 := categoryProductModel.IsCategoryProductExists(categoryProductId3)
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
