package service

import (
	"context"
	"os"
	"testing"

	"github.com/LordRahl90/auth-manager-ot/domain/users/entities"
	"github.com/LordRahl90/auth-manager-ot/domain/users/mocks"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	// this was broken down to make sure that we can compose and decompose items
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreate(t *testing.T) {
	repo := &mocks.MockUserRepo{
		CreateFunc: func(ctx context.Context, user *entities.User) error {
			user.ID = uuid.NewString()
			return nil
		},
	}
	svc := NewUserService(repo)
	user := &entities.User{
		Email:     gofakeit.Email(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
	}
	err := svc.Create(context.Background(), user)
	require.NoError(t, err)
	assert.NotEmpty(t, user.ID)
}

func TestCreateWithError(t *testing.T) {
	repo := &mocks.MockUserRepo{
		CreateFunc: func(ctx context.Context, user *entities.User) error {
			return gorm.ErrInvalidDB
		},
	}
	svc := NewUserService(repo)
	user := &entities.User{
		Email:     gofakeit.Email(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
	}
	err := svc.Create(context.Background(), user)
	assert.EqualError(t, err, gorm.ErrInvalidDB.Error())
}
