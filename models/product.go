package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null; size:50; index"`
	Description string `json:"description" gorm:"not null; size:255"`
	Image       string `json:"image" gorm:"not null; type:varchar(150); index"`
	Quantity    uint   `json:"quantity" gorm:"not null"`
}
