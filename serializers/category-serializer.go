package serializers

import (
	"bookshop/models"
)

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategorySerializer struct {
	Category model.Category
}

// Receiver function for category response
func (serializer *CategorySerializer) Response() CategoryResponse {
	return CategoryResponse{
		ID:   serializer.Category.ID,
		Name: serializer.Category.Name,
	}
}

type CategoriesSerializer struct {
	Categories []model.Category
}

// Receiver function for categoriesSerializers
func (serializer *CategoriesSerializer) Response() []CategoryResponse {
	response := []CategoryResponse{}
	for _, category := range serializer.Categories {
		categoriesSerializer := CategorySerializer{Category: category}
		response = append(response, categoriesSerializer.Response())
	}
	return response
}
