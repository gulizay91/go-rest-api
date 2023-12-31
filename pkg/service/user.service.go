package service

import (
	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/gulizay91/go-rest-api/pkg/repository"
	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	Insert(userEntity entities.User) (*models.ServiceResponseModel, error)
	Get() (*models.ServiceResponseModel, error)
	Delete(id primitive.ObjectID) (bool, error)
}

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(Repo repository.UserRepository) UserService {
	return UserService{Repo: Repo}
}

func (s UserService) Insert(userEntity entities.User) (*models.ServiceResponseModel, error) {
	var newUser models.UserModel
	var res models.ServiceResponseModel = models.ServiceResponseModel{
		Data:    &newUser,
		Success: false,
	}
	if len(userEntity.FirstName) <= 2 {
		res.Message = "FirstName must be valid!"
		return &res, nil
	}

	userEntity.Id = primitive.NewObjectID()
	result, err := s.Repo.Insert(userEntity)

	if err != nil || result == false {
		res.Message = err.Error()
		return &res, err
	}

	res = models.ServiceResponseModel{
		Data: models.UserModel{
			Id:          userEntity.Id.String(),
			SubId:       userEntity.SubId,
			FirstName:   userEntity.FirstName,
			LastName:    userEntity.LastName,
			Email:       userEntity.Email,
			PhoneNumber: userEntity.PhoneNumber,
			BirthDate:   userEntity.BirthDate.Time().String(),
		},
		Success: true,
	}
	return &res, nil
}

func (s UserService) Get(subId string) (*models.ServiceResponseModel, error) {
	var res models.ServiceResponseModel = models.ServiceResponseModel{
		Data:       nil,
		Success:    false,
		StatusCode: "500",
	}
	result, err := s.Repo.Get(subId)
	if err != nil {
		res.Message = err.Error()
		return &res, err
	}

	if result.Id == primitive.NilObjectID {
		res.Message = "mongo: no documents in result"
		res.StatusCode = "404"
		return &res, nil
	}

	res = models.ServiceResponseModel{
		Data: models.UserModel{
			Id:          result.Id.String(),
			SubId:       result.SubId,
			FirstName:   result.FirstName,
			LastName:    result.LastName,
			Email:       result.Email,
			PhoneNumber: result.PhoneNumber,
			BirthDate:   result.BirthDate.Time().String(),
		},
		Success:    true,
		StatusCode: "200",
	}
	return &res, nil
}

func (s UserService) Delete(id primitive.ObjectID) (bool, error) {
	result, err := s.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}
