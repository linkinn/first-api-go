package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutpu struct {
	AccessToken string `json:"access_token"`
}

type Error struct {
	Message string `json:"error"`
}
