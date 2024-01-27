package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id          primitive.ObjectID  `json:"_id"`
	SubId       string              `json:"subId" validate:"required"`
	CreatedDate primitive.DateTime  `json:"createdDate"`
	UpdatedDate primitive.DateTime  `json:"updatedDate"`
	FirstName   string              `json:"firstName" validate:"required,max=50"`
	LastName    string              `json:"lastName" validate:"required,max=50"`
	Email       string              `json:"email" validate:"required,email"`
	PhoneNumber *string             `json:"phoneNumber,omitempty"`
	BirthDate   *primitive.DateTime `json:"birthDate,omitempty"`
	Gender      *string             `json:"gender,omitempty"`
	Media       *Media              `json:"media,omitempty"`
}

type Media struct {
	Images []string `json:"images" validate:"required,gt=0"`
}
