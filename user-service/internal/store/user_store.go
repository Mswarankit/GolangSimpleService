package store

import "github.com/Mswarankit/user-service/internal/models"

type UserStore interface {
	Set(*models.User) error
	Get(id string) (*models.User, error)
	List() []*models.User
}
