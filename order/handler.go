package order

import (
	"ecommerce-api/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

type Handler struct {
	Store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{Store: store}
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var newOrder models.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	if err := h.Store.db.Create(&newOrder).Error; err != nil {
		c.JSON(500, gin.H{"error": "Record not created"})
		return
	}
	c.JSON(201, newOrder)
}

func (h *Handler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	NumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid input format"})
		return
	}
	var order models.Order
	if err := h.Store.db.Where("id = ?", NumId).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusFound, order)
}

func (h *Handler) DeleteOrderByID(c *gin.Context) {
	id := c.Param("id")
	NumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(406, gin.H{"error": "Invalid input format"})
		return
	}
	if err := h.Store.db.Unscoped().Delete(&models.Order{}, NumId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "record Not found"})
		}
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Record has been deleted"})
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	NumId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(406, gin.H{"error": "Invalid input format"})
		return
	}

	var order models.Order
	if err := h.Store.db.Where("id = ?", NumId).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data provided for update"})
		return
	}

	if err := h.Store.db.Model(&models.Order{}).Where("id = ?", NumId).Updates(data).Error; err != nil {
		fmt.Printf("Error while updating order : %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
