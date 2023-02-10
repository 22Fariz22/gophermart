package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"log"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	_, err := r.Pool.Exec(context.Background(),
		"CREATE TABLE if not exists user(ID SERIAL PRIMARY KEY,login TEXT,password TEXT,"+
			"balance_total integer, withdraw_total integer;")
	if err != nil {
		log.Printf("Unable to create table: %v\n", err)
		return err
	}

	fmt.Println("user from db create user:", user)

	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {

	return nil, nil
}
