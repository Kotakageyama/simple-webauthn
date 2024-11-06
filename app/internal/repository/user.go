package repository

import (
	"app/internal/domain"
	"golang.org/x/xerrors"
)

type UserRepository interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	Get(id []byte) (*domain.User, error)
	GetByUserId(userId []byte) (*domain.User, error)
	GetByWorldIDNullifier(nullifier string) (*domain.User, error)
}

type userRepository struct {
	UsersByID        map[string]*domain.User // key: string(user.ID)
	UsersByNullifier map[string]*domain.User // key: World ID nullifier
}

func NewUserRepository() UserRepository {
	return &userRepository{
		UsersByID:        make(map[string]*domain.User),
		UsersByNullifier: make(map[string]*domain.User),
	}
}

func (r *userRepository) Create(user *domain.User) error {
	if _, exists := r.UsersByID[string(user.ID)]; exists {
		return xerrors.New("user already exists")
	}

	r.UsersByID[string(user.ID)] = user

	// If this is a World ID user, also store by nullifier
	if user.AuthMethod == "worldid" {
		r.UsersByNullifier[string(user.ID)] = user // For World ID users, ID is the nullifier
	}

	return nil
}

func (r *userRepository) Update(user *domain.User) error {
	if _, exists := r.UsersByID[string(user.ID)]; !exists {
		return xerrors.New("user not found")
	}

	r.UsersByID[string(user.ID)] = user

	// Update nullifier index if this is a World ID user
	if user.AuthMethod == "worldid" {
		r.UsersByNullifier[string(user.ID)] = user
	}

	return nil
}

func (r *userRepository) Get(id []byte) (*domain.User, error) {
	u, ok := r.UsersByID[string(id)]
	if !ok {
		return nil, xerrors.New("user not found")
	}
	return u, nil
}

func (r *userRepository) GetByUserId(userID []byte) (*domain.User, error) {
	return r.Get(userID)
}

func (r *userRepository) GetByWorldIDNullifier(nullifier string) (*domain.User, error) {
	user, ok := r.UsersByNullifier[nullifier]
	if !ok {
		return nil, xerrors.New("user not found")
	}
	return user, nil
}
