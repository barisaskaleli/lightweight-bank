package request

type RegisterRequest struct {
	Name     string  `json:"name" validate:"required"`
	Surname  string  `json:"surname" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,gte=4,lte=20"`
	Balance  float64 `json:"balance" validate:"required,gt=0"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
}
