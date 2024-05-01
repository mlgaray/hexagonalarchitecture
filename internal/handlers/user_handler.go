package handlers

import (
	"encoding/json"
	"github.com/mlgaray/hexagonalarchitecture/internal/core/ports/user"
	"net/http"
)

type UserHandler struct {
	userServicePort user.UserServicePort
}

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := u.userServicePort.GetAll()
	json.NewEncoder(w).Encode(users)
}

func (u *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	/*	var t models.User

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {

		}
		err = u.userServicePort.SignUp(t)

		if err != nil {

		}*/
	//json.NewEncoder(w).Encode(users)
}

func NewUserHandler(userServicePort user.UserServicePort) *UserHandler {
	return &UserHandler{
		userServicePort: userServicePort,
	}
}
