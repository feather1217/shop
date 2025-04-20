package models

type LineMessageRequest struct {
	UserID  string `json:"userId"`
	Message string `json:"message"`
}