package appuser

type RegisterUserInputDto struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
}

type ResetPasswordDto struct {
	Username string `json:"username" binding:"required"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
