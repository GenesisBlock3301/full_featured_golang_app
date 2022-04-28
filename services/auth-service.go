package services

import (
	"bookshop/config"
	model "bookshop/models"
	"bookshop/validinput"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func RegisterService(userInput validinput.User) (model.User, error) {
	user := model.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	}
	if err := config.DB.Where("name = ?", userInput.Name).First(&user).Error; err == nil {
		return user, errors.New("User already exits")
	}
	user.Password = hashAndSalt([]byte(user.Password))
	config.DB.Create(&user)
	return user, nil
}

func VerifiyCredentialService(email string, password string) (bool, uint) {

	user, err := FindByEmail(email)
	if err != nil {
		return false,0
	}
	return comparePassword([]byte(user.Password),[]byte(password)),uint(user.ID)
	
}


// Get user by id

func GetUserById(id uint) (model.User,error){
	var user model.User
	if err := config.DB.First(&user,id).Error;err !=nil{
		return user,errors.New("User not found")
	}
	return user,nil
}


// Make password hashed
func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(hash)
}

// Find user by email
func FindByEmail(email string) (model.User, error) {
	user := model.User{}
	if err := config.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Compare User password and input password
func comparePassword(hashedPass []byte,inputPass []byte) bool{
	err := bcrypt.CompareHashAndPassword(hashedPass,inputPass)
	if err != nil{
		return false
	}
	return true
}