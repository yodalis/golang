package entity //Pacote com as regras de negócio

import (
	"github.com/yodalis/golang/9-apis/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// VO - value object (objetos de valor) => objeto que representa aquilo que a gente ta fazendo
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"Name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // vai encriptar nossa senha e o segundo parametro é o quanto de capacidade computacional que vai utilizar pra fazer isso
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

// Aqui vamos comparar a senha que veio com a que temos no hash pra ver se a correta
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
