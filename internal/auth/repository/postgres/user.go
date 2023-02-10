package postgres

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/postgres"
)

type User struct {
	ID       uint32 `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) CreateUser(ctx context.Context, user *entity.User) error {

	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {

	return nil, nil
}
