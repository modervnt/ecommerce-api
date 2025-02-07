package product

import (
	"ecommerce-api/models"

	"gorm.io/gorm"
)

/*GetProduct*/
/*CreateProduct*/

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateProduct(product *models.Product) error {
	return s.db.Create(product).Error
}
