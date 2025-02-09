/*
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
		DB.AutoMigrate(&models.Order{})
	}
*/
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
		fmt.Println("Erreur lors de la connexion à la base de données :", err)
		return
	}

	fmt.Println("Connexion à la base de données réussie.")

	// Active les logs SQL pour débogage
	DB = DB.Debug()

	// Migration automatique des tables
	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		fmt.Println("Erreur lors de la migration des tables :", err)
		return
	}

	fmt.Println("Migration des tables réussie.")
}
