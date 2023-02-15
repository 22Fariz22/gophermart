package auth

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

//interface

type UserRepository interface {
	CreateUser(ctx context.Context, l logger.Interface, user *entity.User) error
	GetUser(ctx context.Context, l logger.Interface, username, password string) (*entity.User, error)
}
