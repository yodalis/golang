package dto

// DTO ou Data Transfer é um objeto que armazena dados puros para ser utilizados para
// trafegar dados de forma pura,
// ou seja, sem regra de negócio nenhuma aplicada ao dado

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"access_token"`
}
