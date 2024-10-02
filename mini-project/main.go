package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mini-project-mongo/config"
	"mini-project-mongo/delivery/http"
	domain "mini-project-mongo/domain/usecase"
	"mini-project-mongo/repository"
	"mini-project-mongo/usecase"
)

var (
	server      *gin.Engine
	us          domain.UserUsecaseI
	uc          http.UserController
	ctx         context.Context
	mongoClient *mongo.Client
)

func init() {
	ctx = context.TODO()

	// Mongo
	mongoCon, err := config.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("mongo connection established")
	ur := repository.NewUserRepository(mongoCon)
	us = usecase.NewUserUsecase(ur, ctx)
	uc = http.New(us)
	server = gin.Default()
}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {

		}
	}(mongoClient, ctx)

	basePath := server.Group("/v1")
	uc.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":9090"))

}
