package mappers

import (
	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func MapUserModelToUser(userModel *models.UserModel) *entities.User {
	return &entities.User{
		Id:          MapStringToObjectId(userModel.Id),
		SubId:       userModel.SubId,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
		BirthDate:   primitive.NewDateTimeFromTime(userModel.BirthDate),
		Gender:      string(userModel.Gender),
		Languages:   MapLanguageSliceToStringSlice(userModel.Languages),
		Media:       MapMediaModelToEntity(userModel.Media),
		CreatedDate: MapTimeToDateTimeUtc(userModel.CreatedDate),
		UpdatedDate: MapTimeToDateTime(userModel.UpdatedDate),
	}
}

func MapStringToObjectId(s string) primitive.ObjectID {
	if s == "" {
		return primitive.NilObjectID
	}
	objectId, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return objectId
}

func MapTimeToDateTime(time *time.Time) *primitive.DateTime {
	if time == nil {
		return nil
	}
	tt := primitive.NewDateTimeFromTime(*time)
	return &tt
}

func MapTimeToDateTimeUtc(pTime *time.Time) primitive.DateTime {
	if pTime == nil {
		now := time.Now().UTC()
		return primitive.NewDateTimeFromTime(now)
	}
	return primitive.NewDateTimeFromTime(*pTime)
}

func MapMediaModelToEntity(mediaModel *models.Media) *entities.Media {
	if mediaModel == nil {
		return nil
	}
	return &entities.Media{
		Images: mediaModel.Images,
	}
}

func MapLanguageSliceToStringSlice(languages []*models.Language) []*string {
	if languages == nil {
		return nil
	}
	stringSlice := make([]*string, len(languages))
	for i, language := range languages {
		stringSlice[i] = (*string)(language)
	}
	return stringSlice
}
