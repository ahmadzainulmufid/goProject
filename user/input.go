package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=1,max=8"`
	Occupation string `json:"occupation" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=1,max=8"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
