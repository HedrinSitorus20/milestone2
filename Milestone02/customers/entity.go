package customers

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar" binding:"required"`
}
