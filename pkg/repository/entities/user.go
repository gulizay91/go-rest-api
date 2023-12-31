package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	SubId       string             `bson:"subId,omitempty"`
	FirstName   string             `bson:"firstName,omitempty"`
	LastName    string             `bson:"lastName,omitempty"`
	Email       string             `bson:"email,omitempty"`
	PhoneNumber string             `bson:"phoneNumber,omitempty"`
	BirthDate   primitive.DateTime `bson:"birthDate"`
}