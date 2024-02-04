package services

import (
	"context"
	stdLog "log"
	"time"

	configs "github.com/gulizay91/go-rest-api/config"
	"github.com/gulizay91/go-rest-api/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userRepository repository.UserRepository
var mongoDBClient *mongo.Client

func RegisterRepositories(config *configs.Config) {
	mongoDBClient = connectMongoDB(config)

	userCollection := getCollection(mongoDBClient, config.MongoDB.Database, "users")
	registerUserRepository(userCollection)
}

func connectMongoDB(config *configs.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoDB.Uri))

	if err != nil {
		stdLog.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		stdLog.Fatalln(err)
	}

	stdLog.Printf("Ping mongo database %s", config.MongoDB.Database)
	return client
}

func getCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func ensureIndexes(collection *mongo.Collection, indexes []mongo.IndexModel) error {

	opts := options.CreateIndexes().SetMaxTime(5 * time.Second)
	_, err := collection.Indexes().CreateMany(context.Background(), indexes, opts)
	if err != nil {
		return err
	}
	return nil
}

func registerUserRepository(collection *mongo.Collection) {
	userRepository = repository.NewUserRepository(collection)
	userCollectionIndexes := userRepository.GetCollectionIndexes()
	err := ensureIndexes(collection, userCollectionIndexes)
	if err != nil {
		stdLog.Fatalln(err)
	}
}
