package domain

import (
	"context"
	"mini-project-mongo/domain/models"
)

type UserRepositoryI interface {
	GetAllData(ctx context.Context) (user []models.User, err error)
	InsertData(ctx context.Context, req *models.User) error
	UpdateData(ctx context.Context, req *models.User) error
	DeleteData(ctx context.Context, req *string) error
	GetData(ctx context.Context, username *string) (user *models.User, err error)
}
