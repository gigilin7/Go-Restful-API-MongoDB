package domain

import (
	"context"
	"mini-project-mongo/domain/models"
)

type UserUsecaseI interface {
	CreateUser(ctx context.Context, req *models.User) error
	GetUser(ctx context.Context, req *string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, req *models.User) error
	DeleteUser(ctx context.Context, req *string) error
}
