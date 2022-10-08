package interfaces

import (
	"context"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
)

type UserService interface {
	Create(ctx context.Context, user *entities.User) error
	FindAll(ctx context.Context, page, limnit int) ([]entities.User, error)
	FindOne(ctx context.Context) (*entities.User, error)
}

type UserRepo interface {
	Create(ctx context.Context, user *entities.User) error
	Find(ctx context.Context, id string) (*entities.User, error)
}
