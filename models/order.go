package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  uint    `json:"user_id" gorm:"not null"`
	Total   float64 `json:"total" gorm:"not null"`
	Status  string  `json:"status" gorm:"type:TEXT; default:'pending'; check:status IN ('pending','shipped','delivered','cancelled')"`
	Address string  `json:"address" gorm:"not null; size:255"`
	User    User    `json:"-" gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
