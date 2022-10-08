package mocks

import (
	"context"
	"fmt"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
)

// MockUserRepo user repo mock
type MockUserRepo struct {
	CreateFunc func(ctx context.Context, user *entities.User) error
	FindFunc   func(ctx context.Context, id string) (*entities.User, error)
}

func (m *MockUserRepo) Create(ctx context.Context, user *entities.User) error {
	if m.CreateFunc == nil {
		return fmt.Errorf("create function not initialized")
	}
	return m.CreateFunc(ctx, user)
}

func (m *MockUserRepo) Find(ctx context.Context, id string) (*entities.User, error) {
	if m.FindFunc == nil {
		return nil, fmt.Errorf("find function not initialized")
	}

	return m.FindFunc(ctx, id)
}
