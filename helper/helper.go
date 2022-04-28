package helper

import (
	"bookshop/validinput"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadPostImage(ctx *gin.Context, form *validinput.Post) error {
	fmt.Println("Upload image function calling")
	file, _ := ctx.FormFile("image")
	if file != nil {
		fileName := filepath.Base(file.Filename)
		fmt.Println("FileName", fileName)
		// Upload the file to specific dst.
		if err := ctx.SaveUploadedFile(file, "media/images/"+fileName); err != nil {
			fmt.Println("Upload file error working.")
			return err
		}
		form.Image = "media/images/" + fileName

	}
	return nil
}
