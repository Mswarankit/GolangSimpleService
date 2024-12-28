package store

import (
	"fmt"
	"sync"

	"github.com/Mswarankit/user-service/internal/models"
)

type MemoryStore struct {
	mu    sync.Mutex
	users map[string]*models.User
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]*models.User),
	}
}

func (s *MemoryStore) Set(user *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.ID] = user
	return nil
}

func (s *MemoryStore) Get(id string) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *MemoryStore) List() []*models.User {
	s.mu.Lock()
	defer s.mu.Unlock()
	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}
