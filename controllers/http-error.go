package controllers

type HttpError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Student not found"`
}
