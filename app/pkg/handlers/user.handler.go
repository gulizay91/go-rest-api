package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gulizay91/go-rest-api/pkg/mappers"
	"github.com/gulizay91/go-rest-api/pkg/models"
	"github.com/gulizay91/go-rest-api/pkg/repository/entities"
	"github.com/gulizay91/go-rest-api/pkg/service"
	"github.com/gulizay91/go-rest-api/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{Service: service}
}

// CreateUser
// @Summary create user
// @Description create user
// @Tags users
// @Accept */*
// @Produce json
// @Param   user body  models.UserModel true "User"
// @Success 201 {object} models.ServiceResponseModel
// @Failure 400 {object} models.ServiceResponseModel
// @Failure 409 {object} models.ServiceResponseModel
// @Failure 500 {object} models.ServiceResponseModel
// @Router /api/v1/user [post]
func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var user *models.UserModel
	result := models.NewErrorServiceResponseModel(nil)
	result.StatusCode = http.StatusBadRequest
	if err := c.BodyParser(&user); err != nil {
		result.Message = err.Error()
		return c.Status(result.StatusCode).JSON(result)
	}

	if err := utils.Validate(user); err != nil {
		result.Data = err
		result.Message = "Validation failed"
		return c.Status(result.StatusCode).JSON(result)
	}
	var newUserEntity *entities.User = mappers.MapUserModelToUser(user)
	newUserEntity.CreatedDate = primitive.NewDateTimeFromTime(time.Now().UTC())
	newUserEntity.UpdatedDate = primitive.NewDateTimeFromTime(time.Now().UTC())
	result, err := h.Service.Insert(newUserEntity)

	if err != nil || result.Success == false {
		return c.Status(result.StatusCode).JSON(result)
	}

	return c.Status(http.StatusCreated).JSON(result)
}

// GetUser
// @Summary get user
// @Description get user
// @Tags users
// @Accept */*
// @Produce json
// @Param  subId  path string true  "subId"
// @Success 200 {object} models.ServiceResponseModel
// @Failure 404 {object} models.ServiceResponseModel
// @Failure 500 {object} models.ServiceResponseModel
// @Router /api/v1/users/{subId} [get]
func (h UserHandler) GetUser(c *fiber.Ctx) error {
	subId := c.Params("subId")
	result, err := h.Service.Get(subId)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	if result.Success == false {
		return c.Status(http.StatusNotFound).JSON(result)
	}

	return c.Status(http.StatusOK).JSON(result)
}

// DeleteUser
// @Summary delete user
// @Description delete user
// @Tags users
// @Accept */*
// @Produce json
// @Param  id  path string true  "id"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/user/{id} [delete]
func (h UserHandler) DeleteUser(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.Delete(cnv)

	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Success": true})
}
