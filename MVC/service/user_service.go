package service

import (
	"github.com/yourusername/simple-api/model"
	"sync"
)

type UserService struct {
	users  []model.User
	mu     sync.Mutex
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  []model.User{},
		nextID: 1,
	}
}

func (s *UserService) GetAllUsers() []model.User {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.users
}

func (s *UserService) AddUser(user model.User) model.User {
	s.mu.Lock()
	defer s.mu.Unlock()
	user.ID = s.nextID
	s.nextID++
	s.users = append(s.users, user)
	return user
}
