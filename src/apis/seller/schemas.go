package seller

type CreateSellerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type UpdateSellerRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
