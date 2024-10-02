package usecase

import (
	"context"
	"log"
	"mini-project-mongo/domain/models"
	domain2 "mini-project-mongo/domain/repository"
	domain "mini-project-mongo/domain/usecase"
)

type UserServiceImpl struct {
	userRepo domain2.UserRepositoryI
	ctx      context.Context
}

func (u UserServiceImpl) CreateUser(ctx context.Context, req *models.User) error {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}
	//insert data
	err := u.userRepo.InsertData(ctx, req)
	if err != nil {
		return err
	}
	log.Println("Successfully Inserted Data User")

	return nil
}

func (u UserServiceImpl) GetUser(ctx context.Context, req *string) (*models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.userRepo.GetData(ctx, req)

	return user, err
}

func (u UserServiceImpl) GetAll(ctx context.Context) ([]models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := u.userRepo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show data user with default log")
		return list, err
	}

	return list, err
}

func (u UserServiceImpl) UpdateUser(ctx context.Context, req *models.User) error {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err := u.userRepo.UpdateData(ctx, req)
	if err != nil {
		log.Println("ERROR  : ", err)
		return err
	}

	return nil
}

func (u UserServiceImpl) DeleteUser(ctx context.Context, req *string) error {
	if ctx == nil {
		ctx = context.Background()
	}

	//update data
	err := u.userRepo.DeleteData(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func NewUserUsecase(userRepo domain2.UserRepositoryI, ctx context.Context) domain.UserUsecaseI {
	return &UserServiceImpl{
		userRepo: userRepo,
		ctx:      ctx,
	}
}
