package response

type RegisterResponse struct {
	ID            uint32  `json:"id"`
	AccountNumber string  `json:"account_number"`
	Name          string  `json:"name"`
	Surname       string  `json:"surname"`
	Email         string  `json:"email"`
	Balance       float64 `json:"balance"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
