package database

import (
	"context"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
	"github.com/LordRahl90/auth-manager-ot/domain/users/interfaces"
	"go.opentelemetry.io/otel"

	"github.com/google/uuid"
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
	ctx, span := otel.Tracer("users").Start(ctx, "UserRepo_Create")
	defer span.End()
	m := FromServiceEntity(*user)
	m.ID = uuid.NewString()
	err := u.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		return err
	}
	user.ID = m.ID
	return nil
}

// Find finds a user with an id
func (u *UserRepo) Find(ctx context.Context, id string) (*entities.User, error) {
	return nil, nil
}

// ToServiceEntity return a service entity
func (e *User) ToServiceEntity() entities.User {
	return entities.User{
		ID:        e.ID,
		Email:     e.Email,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
}

// FromServiceEntity return a model from entity
func FromServiceEntity(e entities.User) User {
	return User{
		Email:     e.Email,
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}
}
