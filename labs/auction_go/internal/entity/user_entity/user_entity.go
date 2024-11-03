package user_entity

import (
	"context"

	"github.com/yodalis/golang/labs/auction_go/internal/internal_error"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userID string) (*User, *internal_error.InternalError)
}
