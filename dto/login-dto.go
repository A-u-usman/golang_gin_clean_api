package dto

//LoginDTO is a model that is used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" forrm:"password" binding:"required" validate:"min:6"`
}
