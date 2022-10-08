package users

import (
	"context"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
	"github.com/LordRahl90/auth-manager-ot/domain/users/interfaces"
	"gorm.io/gorm"
)

// UserRepo user repository
type UserRepo struct {
	db *gorm.DB
}

// User db entity for user
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	gorm.Model
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &UserRepo{db: db}
}

// Create repo create
func (u *UserRepo) Create(ctx context.Context, user *entities.User) error {
	return nil
}

// Find finds a user with an id
func (u *UserRepo) Find(ctx context.Context, id string) (*entities.User, error) {
	return nil, nil
}
