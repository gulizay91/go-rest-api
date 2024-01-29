package mappers

import (
	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
)

func MapUserModelToUser(userModel *models.UserModel) *entities.User {
	return &entities.User{
		Id:          userModel.Id,
		SubId:       userModel.SubId,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
		BirthDate:   userModel.BirthDate,
		Gender:      string(userModel.Gender),
		Media:       MapMediaModelToEntity(userModel.Media),
	}
}

func MapMediaModelToEntity(mediaModel *models.Media) *entities.Media {
	if mediaModel == nil {
		return nil
	}
	return &entities.Media{
		Images: mediaModel.Images,
	}
}
