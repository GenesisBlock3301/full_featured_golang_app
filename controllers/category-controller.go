package controller

import (
	model "bookshop/models"
	"bookshop/serializers"
	"bookshop/services"
	"bookshop/validinput"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all categories
func AllCategories(ctx *gin.Context) {
	categories := services.GetAllCategoriesService()
	serializer := serializers.CategoriesSerializer{Categories: categories}
	ctx.JSON(http.StatusOK, serializer.Response())
}

// Insert Category
func InsertCategory(ctx *gin.Context) {
	categoryInput := validinput.Category{}
	err := ctx.ShouldBindJSON(&categoryInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	categoryModel := model.Category{Name: categoryInput.Name}
	category := services.CategoryInsertService(categoryModel)
	serializer := serializers.CategorySerializer{Category: category}
	ctx.JSON(http.StatusCreated, serializer.Response())

}

// Get Category By ID
func CategoryById(ctx *gin.Context) {
	category, err := services.FindCategoryByIdService(ctx.Param("categoryId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	serializer := serializers.CategorySerializer{Category: category}
	ctx.JSON(http.StatusOK, serializer.Response())

}

func CategoryUpdate(ctx *gin.Context) {
	updateInput := validinput.Category{}
	id := ctx.Param("categoryId")
	if err := ctx.ShouldBindJSON(&updateInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	categoryModel := model.Category{
		Name: updateInput.Name}
	category, err := services.CategoryUpdateService(id, categoryModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	serializer := serializers.CategorySerializer{Category: category}
	ctx.JSON(http.StatusOK, serializer.Response())

}

func CategoryDeleteById(ctx *gin.Context) {
	id := ctx.Param("categoryId")
	err := services.CategoryDeleteByIdService(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})
}
