package service

import (
	"context"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
	"github.com/LordRahl90/auth-manager-ot/domain/users/interfaces"
	"github.com/LordRahl90/auth-manager-ot/domain/users/repository/database"
	"go.opentelemetry.io/otel"

	"gorm.io/gorm"
)

type UserService struct {
	repo interfaces.UserRepo
}

// NewUserService create a new user service with database
// any injection here should be another interface
func NewUserService(repo interfaces.UserRepo) interfaces.UserService {
	return &UserService{repo: repo}
}

// DefaultUserService returns a new user service based on the provided db connection
func DefaultUserService(db *gorm.DB) interfaces.UserService {
	repo := database.NewUserRepo(db)
	return NewUserService(repo)
}

func (u *UserService) Create(ctx context.Context, user *entities.User) error {
	ctx, span := otel.Tracer("users").Start(ctx, "UserService_Create")
	defer span.End()
	return u.repo.Create(ctx, user)
}

func (u *UserService) FindAll(ctx context.Context, page, limit int) ([]entities.User, error) {
	return nil, nil
}

func (u *UserService) FindOne(ctx context.Context) (*entities.User, error) {
	return nil, nil
}
