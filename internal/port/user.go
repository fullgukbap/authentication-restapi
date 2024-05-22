package port

import "authentication-restapi/internal/domain"

// InBound (recommend using a handler adapter)
type UserService interface {
	Register(*domain.User) (*domain.User, error)
	Get(uint) (*domain.User, error)
	Update(uint, *domain.User) (*domain.User, error)
	Delete(uint) error
}

// OutBound (recommdn using a persistence adapter)
type UserRepository interface {
	Create(*domain.User) (*domain.User, error)
	Read(uint) (*domain.User, error)
	Udpate(uint, *domain.User) (*domain.User, error)
	Delete(uint) error
}
