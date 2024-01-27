package models

type ValidationErrors struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
