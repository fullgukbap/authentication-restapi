package app

import (
	"authentication-restapi/internal/domain"
	"authentication-restapi/internal/port"
)

type UserService struct {
	UserRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) Register(u *domain.User) (*domain.User, error) {
	return s.UserRepository.Create(u)
}

func (s *UserService) Get(id uint) (*domain.User, error) {
	return s.UserRepository.Read(id)
}

func (s *UserService) Update(id uint, u *domain.User) (*domain.User, error) {
	return s.UserRepository.Udpate(id, u)
}

func (s *UserService) Delete(id uint) error {
	return s.UserRepository.Delete(id)
}
