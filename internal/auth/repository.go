package auth

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
)

//interface

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, username, password string) (*entity.User, error)
}
