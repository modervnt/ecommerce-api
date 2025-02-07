package user

import (
	"ecommerce-api/models"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
} // pourquoi on a ca si on a un package DB deja existant?

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByID(id uint) (*models.User, error) { // pourquoi on utilise un pointeur ici?
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (s *Store) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}
