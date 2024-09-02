package repository

import (
	"app/internal/domain"

	"golang.org/x/xerrors"
)

type UserRepository interface {
	Insert(id domain.SessionID, user *domain.User) error
	Get(id domain.SessionID) (*domain.User, error)
	GetByUserId(userId []byte) (*domain.User, error)
}

type userRepository struct {
	Users map[domain.SessionID]*domain.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		Users: map[domain.SessionID]*domain.User{},
	}
}

func (r *userRepository) Insert(id domain.SessionID, user *domain.User) error {
	if _, exists := r.Users[id]; exists {
		return xerrors.New("user already exists")
	}
	r.Users[id] = user
	return nil
}

func (r *userRepository) Get(id domain.SessionID) (*domain.User, error) {
	u, ok := r.Users[id]
	if !ok {
		return nil, xerrors.New("user not found")
	}
	return u, nil
}

func (r *userRepository) GetByUserId(userID []byte) (*domain.User, error) {
	for _, u := range r.Users {
		if string(u.ID) == string(userID) {
			return u, nil
		}
	}
	return nil, xerrors.New("user not found")
}
