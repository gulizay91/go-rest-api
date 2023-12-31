package repository

import (
	"context"
	"log"
	"time"

	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	Insert(user entities.User) (bool, error)
	GetAll() ([]entities.User, error)
	Get(subId string) (entities.User, error)
	Delete(id primitive.ObjectID) (bool, error)
}

type UserRepository struct {
	Users *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return UserRepository{Users: collection}
}

func (r UserRepository) Insert(user entities.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if user.Id == primitive.NilObjectID {
		user.Id = primitive.NewObjectID()
	}

	result, err := r.Users.InsertOne(ctx, user)

	if result.InsertedID == nil || err != nil {
		return false, err
	}
	return true, nil
}

func (r UserRepository) GetAll() ([]entities.User, error) {
	var user entities.User
	var users []entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.Users.Find(ctx, bson.M{})

	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&user); err != nil {
			log.Panicln(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r UserRepository) Get(subId string) (entities.User, error) {
	var user entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$and": []bson.M{
			{"subId": subId},
		},
	}

	if err := r.Users.FindOne(ctx, filter).Decode(&user); err != nil {

		if err.Error() == "mongo: no documents in result" {
			return user, nil
		}

		log.Panicln(err)
		return user, err
	}

	return user, nil
}

func (r UserRepository) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.Users.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}
