package validinput

type Post struct{
	Title string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	CategoryId uint `form:"category" json:"category" binding:"required"`
	Image string `form:"-" json:"-" binding:"omitempty"`
}
