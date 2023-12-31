package services

import (
	"context"
	"log"
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
	userRepository = repository.NewUserRepository(userCollection)
}

func connectMongoDB(config *configs.Config) *mongo.Client {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoDB.Uri).SetTimeout(20*time.Second))

	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Ping %s", config.MongoDB.Database)
	return client
}

func getCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}