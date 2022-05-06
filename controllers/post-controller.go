package controller

import (
	"bookshop/helper"
	model "bookshop/models"
	"bookshop/serializers"
	"bookshop/services"
	"bookshop/validinput"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get all posts
func AllPosts(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	postCount, posts := services.GetAllPostService(limit, offset)
	serializer := serializers.PostsSerializer{Posts: posts}
	ctx.JSON(http.StatusOK, gin.H{
		"totalPost": postCount,
		"posts":     serializer.Response(),
	})
}

// Looking for post by ID
func FindByPostId(ctx *gin.Context) {
	id := ctx.Param("postId")
	post ,err := services.PostFindById(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Id not found."})
	}
	serializer := serializers.PostSerializer{Post: post}
	ctx.JSON(http.StatusOK,serializer.Response())
}

// Insert post
func InsertPost(ctx *gin.Context) {
	form := validinput.Post{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := helper.UploadPostImage(ctx, &form); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	postModel := model.Post{
		Title:       form.Title,
		Description: form.Description,
		Image:       form.Image,
		CategoryId:  form.CategoryId,
	}
	post, err := services.InsertPostService(postModel)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Post not created",
		})
	}
	serializer := serializers.PostSerializer{Post: post}
	ctx.JSON(http.StatusCreated, serializer.Response())
}

// Update post
func UpdatePost(ctx *gin.Context) {
	updatePost := validinput.Post{}
	postId := ctx.Param("postId")
	if err := ctx.ShouldBind(&updatePost); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Upload updated image
	if err := helper.UploadPostImage(ctx, &updatePost); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	post, err := services.PostUpdateService(postId, updatePost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	serializer := serializers.PostSerializer{Post: post}
	ctx.JSON(http.StatusOK, serializer.Response())
}

// Upload post image
func UploadPostImage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Upload post image",
	})
}

// Delete by id
func DeleteByPostId(ctx *gin.Context) {
	postId := ctx.Param("postId")
	err := services.PostDeletedByIdService(postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted.",
	})
}
