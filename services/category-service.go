package services

import (
	"bookshop/config"
	model "bookshop/models"
)

// Get all categories
func GetAllCategoriesService() []model.Category {
	var categories []model.Category
	config.DB.Find(&categories)
	return categories
}

// Insert category
func CategoryInsertService(category model.Category) model.Category {
	config.DB.Create(&category)
	return category
}

func FindCategoryByIdService(categoryId string) (model.Category, error) {
	var category model.Category
	if err := config.DB.First(&category, categoryId).Error; err != nil {
		return category, err
	}
	return category, nil
}

func CategoryUpdateService(categoryId string,input model.Category) (model.Category, error) {
	var category model.Category
	if err := config.DB.First(&category, categoryId).Error; err != nil {
		return category, err
	}
	config.DB.First(&category).Update(input)
	return category,nil
}


func CategoryDeleteByIdService(categoryId string) error{
	var category model.Category
	if err := config.DB.First(&category,categoryId).Error;err !=nil{
		return err
	}
	config.DB.Delete(&category)
	return nil
}