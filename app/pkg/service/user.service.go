package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/gulizay91/go-rest-api/pkg/repository"
	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	Insert(userEntity entities.User) (*models.ServiceResponseModel, error)
	Get() (*models.ServiceResponseModel, error)
	Delete(id primitive.ObjectID) (bool, error)
	UpdateMediaImages(subId string, mediaImages []string) (*models.ServiceResponseModel, error)
}

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(Repo repository.UserRepository) UserService {
	return UserService{Repo: Repo}
}

func (s UserService) Insert(userEntity *entities.User) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)
	if len(userEntity.FirstName) <= 2 {
		res.Message = "FirstName must be valid!"
		res.StatusCode = http.StatusBadRequest
		return &res, nil
	}

	userEntity.Id = primitive.NewObjectID()
	result, err := s.Repo.Insert(userEntity)

	if err != nil || result == false {
		res.Message = err.Error()
		return &res, err
	}

	if err != nil || result == false {
		res.Message = err.Error()
		if mongo.IsDuplicateKeyError(err) {
			res.StatusCode = http.StatusConflict
		}
		return &res, err
	}

	res = *models.NewSuccessServiceResponseModel(userEntity)
	res.StatusCode = http.StatusCreated
	return &res, nil
}

func (s UserService) Get(subId string) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)
	result, err := s.Repo.Get(subId)
	if err != nil {
		res.Message = err.Error()
		return &res, err
	}

	if result.Id == primitive.NilObjectID {
		res.Message = "mongo: no documents in result"
		res.StatusCode = http.StatusNotFound
		return &res, nil
	}

	res = *models.NewSuccessServiceResponseModel(result)
	return &res, nil
}

func (s UserService) Delete(id primitive.ObjectID) (bool, error) {
	result, err := s.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func (s UserService) UpdateMediaImages(subId string, mediaImages []string) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = *models.NewErrorServiceResponseModel(nil)

	if len(mediaImages) < 2 || len(mediaImages) > 6 {
		res.Message = "Media Images len must be between 2 and 6!"
		res.StatusCode = http.StatusBadRequest
		return &res, nil
	}
	keyValue := bson.D{{"media.images", mediaImages}}
	result, err := s.Repo.UpdateOne(subId, keyValue)

	if err != nil && result.Id == primitive.NilObjectID {
		res.Message = err.Error()
		res.StatusCode = http.StatusNoContent
		return &res, err
	}

	if err != nil {
		res.Message = err.Error()
		return &res, err
	}

	res = *models.NewSuccessServiceResponseModel(result)
	res.StatusCode = http.StatusOK
	return &res, nil
}
