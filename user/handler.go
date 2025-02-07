package user

import (
	"ecommerce-api/auth"
	"ecommerce-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Store *Store //d'ou viens le store // le store viens du store.go //ou est la valeur de retour
}

func NewHandler(store *Store) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) GetUser(c *gin.Context) { // qu'est ce que Gin.context?
	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	user, err := h.Store.GetUserByID(uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		fmt.Errorf("error of hashing %v", err)
	}
	newUser.Password = hashedPassword

	if err := h.Store.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) LoginUser(c *gin.Context) {

	var loginData models.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, err := h.Store.GetUserByEmail(loginData.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// compare the password

	if !auth.ComparePassWord(user.Password, []byte(loginData.Password)) {
		c.JSON(http.StatusOK, gin.H{"connection": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

	//then add jwt authentification
}
