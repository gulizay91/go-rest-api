package models

import (
	"time"
)

// swagger:parameters UserModel
type UserModel struct {
	Id          string      `json:"-"`
	SubId       string      `json:"subId" validate:"required"`
	CreatedDate *time.Time  `json:"createdDate,omitempty"`
	UpdatedDate *time.Time  `json:"updatedDate,omitempty"`
	FirstName   string      `json:"firstName" validate:"required,max=50"`
	LastName    string      `json:"lastName" validate:"required,max=50"`
	Email       string      `json:"email" validate:"required,email"`
	PhoneNumber *string     `json:"phoneNumber,omitempty"`
	BirthDate   time.Time   `json:"birthDate" validate:"required"`
	Gender      Gender      `json:"gender" validate:"enum" example:"female"`
	Languages   []*Language `json:"languages,omitempty" validate:"dive,enum"`
	Media       *Media      `json:"media,omitempty"`
}

type Media struct {
	Images []string `json:"images" validate:"required,gt=0"`
}

func NewUserModel(subId string, createdDate *time.Time) *UserModel {
	return &UserModel{
		SubId:       subId,
		CreatedDate: createdDate,
	}
}
