package models

type User struct {
	ID         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required,min=2"`
	SignupTime int64  `json:"signupTime" binding:"required,min=1850"`
}
