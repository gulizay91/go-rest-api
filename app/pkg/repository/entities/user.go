package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID  `bson:"_id"`
	SubId       string              `bson:"subId"`
	CreatedDate primitive.DateTime  `bson:"createdDate"`
	UpdatedDate *primitive.DateTime `bson:"updatedDate,omitempty"`
	FirstName   string              `bson:"firstName"`
	LastName    string              `bson:"lastName"`
	Email       string              `bson:"email"`
	PhoneNumber *string             `bson:"phoneNumber,omitempty"`
	BirthDate   primitive.DateTime  `bson:"birthDate"`
	Gender      string              `bson:"gender"`
	Languages   []*string           `bson:"languages"`
	Media       *Media              `bson:"media,omitempty"`
}

type Media struct {
	Images []string `bson:"images"`
}
