package handlers

import (
	"net/http"
	"user-service/internal/store"

	"github.com/Mswarankit/user-service/internal/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	store store.UserStore
}

func NewUserHandler(store store.UserStore) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.Set(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) ListUser(c *gin.Context) {
	users := h.store.List()
	c.JSON(http.StatusOK, users)
}
