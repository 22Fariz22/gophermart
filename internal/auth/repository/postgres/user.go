package postgres

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	_, err := r.Pool.Exec(ctx, "INSERT INTO users(login, password) values($1, $2);", user.Login, user.Password)
	if err != nil {
		log.Print("err (1) in db CreateUser: ", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {

	return nil, nil
}
