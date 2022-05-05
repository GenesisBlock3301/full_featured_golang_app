package controller

import (
	"bookshop/serializers"
	"bookshop/services"
	"bookshop/validinput"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var userInput validinput.User
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := services.RegisterService(userInput)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, user)
}

func Login(ctx *gin.Context) {
	var userInput validinput.User
	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValidCredential, userId := services.VerifiyCredentialService(userInput.Email, userInput.Password)
	if isValidCredential {
		token, refresh, err := services.GenerateTokenPair(userId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credential"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token":         token,
			"refresh_token": refresh,
		})
	}
}

// Get Current User
func CurrentUser(ctx *gin.Context) {
	user_id, err := services.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := services.GetUserById(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	serializer := serializers.UserSerializer{User: user}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": serializer.Response()})

}
