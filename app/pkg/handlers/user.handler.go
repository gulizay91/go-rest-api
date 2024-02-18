package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	configs "github.com/gulizay91/go-rest-api/config"
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
	UserService service.UserService
	AwsService  service.AwsService
	AppConfig   *configs.Config
}

func NewUserHandler(userService service.UserService, awsService service.AwsService, config *configs.Config) UserHandler {
	return UserHandler{UserService: userService, AwsService: awsService, AppConfig: config}
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

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		result.Message = err.Error()
		log.Error(err)
		return c.Status(result.StatusCode).JSON(result)
	}

	if err := utils.Validate(user); err != nil {
		result.Data = err
		result.Message = "Validation failed"
		return c.Status(result.StatusCode).JSON(result)
	}
	var newUserEntity *entities.User = mappers.MapUserModelToUser(user)
	newUserEntity.CreatedDate = primitive.NewDateTimeFromTime(time.Now().UTC())
	result, err := h.UserService.Insert(newUserEntity)

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
// @Router /api/v1/user/{subId} [get]
func (h UserHandler) GetUser(c *fiber.Ctx) error {
	subId := c.Params("subId")
	result, err := h.UserService.Get(subId)

	if err != nil {
		log.Error(err)
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

	result, err := h.UserService.Delete(cnv)

	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Success": true})
}

// UploadUserImages
// @Summary upload images
// @Description upload user's images
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param  subId  path string true  "subId"
// @Success 200 {object} models.ServiceResponseModel
// @Failure 404 {object} models.ServiceResponseModel
// @Failure 400 {object} models.ServiceResponseModel
// @Failure 500 {object} models.ServiceResponseModel
// @Router /api/v1/user/{subId}/upload-media [patch]
func (h UserHandler) UploadUserImages(c *fiber.Ctx) error {
	subId := c.Params("subId")
	result := models.NewErrorServiceResponseModel(nil)
	result.StatusCode = http.StatusBadRequest

	// Parse the multipart form:
	form, err := c.MultipartForm()
	if err != nil {
		result.Message = err.Error()
		log.Error(err)
		return c.Status(result.StatusCode).JSON(result)
	}

	// Get all files from "images" key:
	files := form.File["images"] // => []*multipart.FileHeader
	if len(files) > 6 {
		result.Message = "The number of objects cannot be greater than 6!"
		return c.Status(result.StatusCode).JSON(result)
	}

	// get user from db
	resultExistUser, err := h.UserService.Get(subId)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	if resultExistUser.Success == false {
		return c.Status(http.StatusNotFound).JSON(resultExistUser)
	}

	// upload images to s3
	var uploadModel *models.UploadS3FileModel = &models.UploadS3FileModel{
		BucketName:             h.AppConfig.AwsService.S3Service.Bucket,
		FilePath:               models.UserMediaPath + "/" + subId,
		Files:                  files,
		BeforeDeleteAllObjects: true,
		CdnUrl:                 h.AppConfig.AwsService.S3Service.CdnUrl,
	}
	result, err = h.AwsService.UploadToS3Path(uploadModel)
	if err != nil || result.Success == false {
		log.Error(err)
		return c.Status(result.StatusCode).JSON(result)
	}

	// update user media urls
	resultData := result.Data.(models.UploadedS3FileModel)
	var imageUrls []string
	if resultData.FileCdnUrls != nil {
		imageUrls = resultData.FileCdnUrls
	} else {
		imageUrls = resultData.FileS3Urls
	}
	resultUpdate, err := h.UserService.UpdateMediaImages(subId, imageUrls)

	if err != nil || resultUpdate.Success == false {
		return c.Status(result.StatusCode).JSON(resultUpdate)
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetUserImages
// @Summary get images
// @Description get user's images
// @Tags users
// @Accept json
// @Produce json
// @Param  subId  path string true  "subId"
// @Success 200 {object} models.ServiceResponseModel
// @Failure 404 {object} models.ServiceResponseModel
// @Failure 400 {object} models.ServiceResponseModel
// @Failure 500 {object} models.ServiceResponseModel
// @Router /api/v1/user/{subId}/media [get]
func (h UserHandler) GetUserImages(c *fiber.Ctx) error {
	subId := c.Params("subId")
	result := models.NewErrorServiceResponseModel(nil)
	result.StatusCode = http.StatusBadRequest

	// Get all files from path (media/subId)
	var listModel *models.ListS3FileModel = &models.ListS3FileModel{
		BucketName: h.AppConfig.AwsService.S3Service.Bucket,
		FilePath:   models.UserMediaPath + "/" + subId,
	}
	result, err := h.AwsService.ListObjectsFromS3Path(listModel)
	if err != nil || result.Success == false {
		log.Error(err)
		return c.Status(result.StatusCode).JSON(result)
	}

	return c.Status(http.StatusOK).JSON(result)
}
