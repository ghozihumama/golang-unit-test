package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRespository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRespository}

func TestCatgoryService_GetNotFound(t *testing.T) {

	//program mock
	categoryRespository.Mock.On("FindById", "1").Return(nil)
	cat, err := categoryService.Get("1")

	assert.Nil(t, cat)
	assert.NotNil(t, err)

}

func TestCatgoryService_GetSuccess(t *testing.T) {

	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	//program mock
	categoryRespository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
