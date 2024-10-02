package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mini-project-mongo/domain/models"
	domain "mini-project-mongo/domain/repository"
)

type UserRepository struct {
	mongoDB *mongo.Database
}

func (c UserRepository) GetData(ctx context.Context, username *string) (user *models.User, err error) {
	var users *models.User
	query := bson.D{bson.E{Key: "name", Value: username}}
	errs := c.mongoDB.Collection("students").FindOne(ctx, query).Decode(&users)
	return users, errs
}

func (c UserRepository) GetAllData(ctx context.Context) (userResp []models.User, err error) {
	query, err := c.mongoDB.Collection("students").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return []models.User{}, err
	}
	defer query.Close(ctx)

	listDataUser := make([]models.User, 0)
	for query.Next(ctx) {
		var row models.User
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataUser = append(listDataUser, row)
	}

	return listDataUser, err
}

func (c UserRepository) InsertData(ctx context.Context, req *models.User) error {

	_, err := c.mongoDB.Collection("students").InsertOne(ctx, req)
	if err != nil {
		log.Println("error")
	}

	return err
}

func (c UserRepository) UpdateData(ctx context.Context, user *models.User) error {
	dataUpdateName := bson.M{"name": user.Name}
	dataObjectID := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}, primitive.E{Key: "age", Value: user.Age}, primitive.E{Key: "address", Value: user.Address}}}}
	_, err := c.mongoDB.Collection("students").UpdateOne(ctx, dataUpdateName, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}

func (c UserRepository) DeleteData(ctx context.Context, req *string) error {

	dataUpdateName := bson.M{"name": req}
	_, err := c.mongoDB.Collection("students").DeleteOne(ctx, dataUpdateName)
	if err != nil {
		log.Println("error")
	}

	return err
}

func NewUserRepository(mongo *mongo.Database) domain.UserRepositoryI {
	return &UserRepository{
		mongoDB: mongo,
	}
}
