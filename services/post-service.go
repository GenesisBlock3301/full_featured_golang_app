package services

import (
	"bookshop/config"
	model "bookshop/models"
	"bookshop/validinput"
	"strconv"
)


// @notice Get all posts and post count
func GetAllPostService(limit string, offset string) (int64, []model.Post) {
	var count int64
	config.DB.Model(&model.Post{}).Count(&count)
	postLimit, err := strconv.Atoi(limit)
	if err != nil || postLimit == 0 {
		postLimit = -1
	}
	postOffset, err := strconv.Atoi(offset)
	if err != nil || postOffset == 0 {
		postOffset = -1
	}
	var posts []model.Post
	config.DB.Limit(postLimit).Offset(postOffset).Preload("Category").Find(&posts)
	return count, posts
}

// Insert post
func InsertPostService(post model.Post) (model.Post, error) {
	if err := config.DB.Create(&post).Error; err != nil {
		return post, err
	}
	config.DB.Preload("Category").First(&post, post.ID)
	return post, nil
}

// Update post service
func PostUpdateService(postId string,postInput validinput.Post) (model.Post,error){
	var post model.Post

	// Find post according to post
	post,err := PostFindById(postId)
	if err != nil{
		return post,err
	}
	// Checking if post input contain file or not
	if postInput.Image == ""{
		postInput.Image = post.Image
	}

	// Update and save post model
	config.DB.Model(&post).Update(postInput)
	
	// Collecting post as well as category according to foreignkey
	postWithCategory,err := PostFindByWithCategory(postId)
	if err != nil{
		return postWithCategory,err
	}
	return postWithCategory,nil
}

// Post Delete by ID
func PostDeletedByIdService(postId string) (error){
	post,err := PostFindById(postId)
	if err != nil{
		return err
	}
	config.DB.Delete(&post)
	return nil
}

// Find post according to id
func PostFindById(id string) (model.Post, error) {
	var post model.Post
	if err := config.DB.Preload("Category").First(&post, id).Error; err != nil {
		return post, err
	}
	return post, nil
}

func PostFindByWithCategory(postId string)(model.Post,error){
	var post model.Post
	if err := config.DB.Preload("Category").First(&post,postId).Error;err !=nil{
		return post,err
	}
	return post,nil
}
