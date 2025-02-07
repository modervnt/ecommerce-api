package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"not null; unique"`
	Password  string `json:"password" gorm:"not null"`
}

type UserStore interface {
	CreateUser()
}
