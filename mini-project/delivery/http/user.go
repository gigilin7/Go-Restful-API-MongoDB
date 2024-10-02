package http

import (
	"github.com/gin-gonic/gin"
	responses "mini-project-mongo/delivery/response"
	"mini-project-mongo/domain/models"
	domain "mini-project-mongo/domain/usecase"
	"net/http"
)

type UserController struct {
	UserUsecase domain.UserUsecaseI
}

func New(userservice domain.UserUsecaseI) UserController {
	return UserController{
		UserUsecase: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserUsecase.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	var username = ctx.Param("name")
	user, err := uc.UserUsecase.GetUser(ctx, &username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: user})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserUsecase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: users})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserUsecase.UpdateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var username = ctx.Param("name")
	err := uc.UserUsecase.DeleteUser(ctx, &username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success"})
}
