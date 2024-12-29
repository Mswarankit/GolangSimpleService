package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mswarankit/user-service/internal/models"
	"github.com/Mswarankit/user-service/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *UserHandler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := store.NewMemoryStore()
	handler := NewUserHandler(store)
	return router, handler
}

func TestGetUser(t *testing.T) {
	router, handler := setupTestRouter()
	router.GET("/users/:id", handler.GetUser)

	// First create a user
	user := &models.User{
		ID:         "1",
		Name:       "Virat Kohli",
		SignupTime: 17354681722377,
	}
	handler.store.Set(user)

	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, response.ID)
}
