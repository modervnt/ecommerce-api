package product

import (
	"ecommerce-api/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
	}

	if err := h.Store.CreateProduct(&newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	NumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid input format"})
		return
	}
	theProduct, err := h.Store.GetProductByID(NumId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, theProduct)
}

func (h *Handler) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	NumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Servor problem"})
		return
	}
	err = h.Store.DeleteByID(NumId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Servor problem"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	// Récupère l'ID depuis l'URL
	id := c.Param("id")
	fmt.Printf("ID récupéré : %s\n", id)

	NumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format input"})
		return
	}

	//Verfifie si le produit existe
	var product models.Product
	if err := h.Store.db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	// Récupère les données JSON dans une map
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Affiche les données reçues pour débogage
	fmt.Printf("Données reçues pour mise à jour : %+v\n", data)

	// Vérifie que des données ont été fournies
	if len(data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data provided for update"})
		return
	}

	// Effectue la mise à jour partielle
	if err := h.Store.db.Model(&models.Product{}).Where("id = ?", NumId).Updates(data).Error; err != nil {
		fmt.Printf("Erreur lors de la mise à jour du produit : %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Renvoie une réponse de succès
	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
