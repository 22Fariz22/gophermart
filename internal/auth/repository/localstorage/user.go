package localstorage

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/entity"
	"sync"
)

type UserLocalStorage struct {
	users map[uint32]*entity.User
	mu    *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[uint32]*entity.User),
		mu:    new(sync.Mutex),
	}
}
func (s *UserLocalStorage) CreateUser(ctx context.Context, user *entity.User) error {
	s.mu.Lock()
	s.users[user.ID] = user
	s.mu.Unlock()

	return nil
}

func (s *UserLocalStorage) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.Login == username && user.Password == password {
			return user, nil
		}
	}
	return nil, auth.ErrUserNotFound
}
