package db

import (
	"ecommerce-api/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erreur lors de la connexion a la base de donnees :", err)
		return
	}

	//Migration automatique des tables
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
