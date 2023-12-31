package models

type UserModel struct {
	Id          string `json:"id,omitempty"`
	SubId       string `json:"subId,omitempty"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	BirthDate   string `json:"birthDate"`
}