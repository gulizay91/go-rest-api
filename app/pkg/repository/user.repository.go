package repository

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"time"

	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserRepository interface {
	GetCollectionIndexes() []mongo.IndexModel
	Insert(user entities.User) (bool, error)
	GetAll() ([]entities.User, error)
	Get(subId string) (entities.User, error)
	Delete(id primitive.ObjectID) (bool, error)
	ReplaceOne(subId string, user *entities.User) (bool, error)
	UpdateOne(subId string, setKeyValue bson.D) (entities.User, error)
}

type UserRepository struct {
	Users *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return UserRepository{Users: collection}
}

func (r UserRepository) GetCollectionIndexes() []mongo.IndexModel {
	var indexes []mongo.IndexModel
	// compound unique index
	indexes = append(indexes, mongo.IndexModel{
		Keys:    bson.D{{Key: "subId", Value: 1}, {Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	// index
	indexes = append(indexes, mongo.IndexModel{
		Keys:    bson.D{{Key: "subId", Value: 1}},
		Options: options.Index(),
	})
	return indexes
}

func (r UserRepository) Insert(user *entities.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if user.Id == primitive.NilObjectID {
		user.Id = primitive.NewObjectID()
	}

	result, err := r.Users.InsertOne(ctx, user)

	if err != nil || result.InsertedID == nil {
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
		log.Panic(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&user); err != nil {
			log.Panic(err)
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

		log.Panic(err)
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

func (r UserRepository) ReplaceOne(subId string, user *entities.User) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$and": []bson.M{
			{"subId": subId},
		},
	}

	result, err := r.Users.ReplaceOne(ctx, filter, user)

	if err != nil {
		log.Panic(err)
		return false, err
	}

	return result.ModifiedCount == result.MatchedCount, nil
}

func (r UserRepository) UpdateOne(subId string, setKeyValue bson.D) (entities.User, error) {
	var user entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$and": []bson.M{
			{"subId": subId},
		},
	}

	update := bson.D{{"$set", setKeyValue}}

	result, err := r.Users.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Panic(err)
		return user, err
	}

	if result.ModifiedCount == result.MatchedCount {
		if err := r.Users.FindOne(ctx, filter).Decode(&user); err != nil {

			if err.Error() == "mongo: no documents in result" {
				return user, nil
			}

			log.Panic(err)
			return user, err
		}
	}

	return user, nil
}
