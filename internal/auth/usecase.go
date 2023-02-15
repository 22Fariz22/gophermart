package auth

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, l logger.Interface, username, password string) error
	SignIn(ctx context.Context, l logger.Interface, username, password string) (string, error)
	ParseToken(ctx context.Context, l logger.Interface, accessToken string) (*entity.User, error)
}
