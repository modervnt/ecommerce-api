package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string  `json:"firstName" gorm:"not null; size:50"`
	LastName  string  `json:"lastName" gorm:"not null;size:50"`
	Email     string  `json:"email" gorm:"not null; unique;type:varchar(100)"`
	Password  string  `json:"password" gorm:"not null;size:100"`
	Orders    []Order `json:"-" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
