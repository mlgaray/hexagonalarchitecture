package services

import (
	"github.com/mlgaray/hexagonalarchitecture/internal/core/models"
	"github.com/mlgaray/hexagonalarchitecture/internal/core/ports/user"
)

type UserService struct {
	repo user.UserRepositoryPort
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) SignUp(user models.User) error {
	return s.repo.SignUp(user)
}

func NewUserService(repository user.UserRepositoryPort) *UserService {
	return &UserService{repo: repository}
}
