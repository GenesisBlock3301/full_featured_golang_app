package validinput


type Category struct{
	Name string `json:"name" binding:"required,max=100"`
}