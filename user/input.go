package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Job      string `json:"job" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
