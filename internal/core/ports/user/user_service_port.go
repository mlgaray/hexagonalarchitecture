package user

import (
	"github.com/mlgaray/hexagonalarchitecture/internal/core/models"
)

type UserServicePort interface {
	GetAll() ([]models.User, error)
	SignUp(models.User) error
}
