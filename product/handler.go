package product

import (
	"ecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid valid payload"})
	}

	if err := h.Store.CreateProduct(&newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}
