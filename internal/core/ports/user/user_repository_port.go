package user

import (
	"github.com/mlgaray/hexagonalarchitecture/internal/core/models"
)

type UserRepositoryPort interface {
	GetAll() ([]models.User, error)
	SignUp(models.User) error
}
