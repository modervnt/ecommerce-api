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

func (s *Store) GetProductByID(id int) (*models.Product, error) {
	var product models.Product
	if err := s.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *Store) DeleteByID(id int) error {
	/*if err := s.db.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}*/
	if err := s.db.Unscoped().Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
