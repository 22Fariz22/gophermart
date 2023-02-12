package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/rs/zerolog/log"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	_, err := r.Pool.Exec(ctx, "INSERT INTO users(login, password) values($1, $2);", user.Login, user.Password)
	// не различает есть ли такой же или просто нету такого вообще, оба варианта это err.а нужны 409 и 500 статусы
	if err != nil {
		log.Print("err (1) in db CreateUser: ", err)
		return err
	}
	//r.Close()
	return nil
}

func (r *UserRepository) GetUser(ctx context.Context, login, password string) (*entity.User, error) {
	row, err := r.Pool.Query(ctx, "select user_id,login,password from users where login = $1 and password = $2", login, password)
	if err != nil {
		log.Print("db-getuser()-row.Scan()")
		return nil, err
	}
	defer row.Close()

	rows := make([]User, 1)

	for row.Next() {
		var u User
		fmt.Println(row.Values())
		err := row.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			return nil, auth.ErrUserNotFound
		}

		rows = append(rows, u)
	}
	if len(rows) < 2 {
		return nil, auth.ErrUserNotFound
	}

	fmt.Println("db-GetUser()-rows", rows)
	fmt.Println("db-GetUser()-len(rows)", len(rows))
	fmt.Println("db-GetUser()-rows[0]", rows[0])
	fmt.Println("db-GetUser()-rows[1]", rows[1])
	return toEntity(&rows[1]), nil
}

func toEntity(u *User) *entity.User {
	return &entity.User{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}
