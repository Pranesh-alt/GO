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

func (s *UserService) GetUserByID(id int) (model.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.ID == id {
			return user, true
		}
	}
	return model.User{}, false
}
func (s *UserService) AddUser(user model.User) model.User {
	s.mu.Lock()
	defer s.mu.Unlock()
	user.ID = s.nextID
	s.nextID++
	s.users = append(s.users, user)
	return user
}

func (s *UserService) UpdateUser(id int, updated model.User) (model.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users[i].Name = updated.Name
			s.users[i].Email = updated.Email
			return s.users[i], true
		}
	}
	return model.User{}, false
}

func (s *UserService) DeleteUser(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return true
		}
	}
	return false
}
