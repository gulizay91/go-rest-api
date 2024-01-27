package models

import "net/http"

type ServiceResponseModel struct {
	Data       interface{} `json:"data"`
	Success    bool        `json:"success,omitempty"`
	StatusCode int         `json:"statusCode"`
	ErrorCode  string      `json:"errorCode,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func NewErrorServiceResponseModel(data *interface{}) *ServiceResponseModel {
	return &ServiceResponseModel{
		Data:       data,
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Message:    "Something went wrong!",
	}
}

func NewSuccessServiceResponseModel(data interface{}) *ServiceResponseModel {
	return &ServiceResponseModel{
		Data:       data,
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Success",
	}
}
