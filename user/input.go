package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
<<<<<<< HEAD
<<<<<<< HEAD

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
=======
>>>>>>> 0c61164ccc255f9a984ef9b8004fed0f398f2c09
=======
>>>>>>> 0c61164ccc255f9a984ef9b8004fed0f398f2c09
