package domain

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Message string `json:"message" gorm:"type:text;not null"`
}
