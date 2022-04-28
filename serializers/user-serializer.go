package serializers

import model "bookshop/models"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserSerializer struct {
	User model.User
}

func (serializer *UserSerializer) Response() UserResponse {
	return UserResponse{
		ID: uint(serializer.User.ID),
		Name: serializer.User.Name,
		Email: serializer.User.Email,
	}
}
