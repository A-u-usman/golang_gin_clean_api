package dto

//RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" `
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" validate:"min:6" binding:"required"`
}

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password,omitempty" form:"password,omitempty"  binding:"required"`
}
