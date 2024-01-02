package models

type ServiceResponseModel struct {
	Data       interface{} `json:"data"`
	Success    bool        `json:"success,omitempty"`
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
}