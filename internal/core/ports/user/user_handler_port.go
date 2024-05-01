package user

import (
	"net/http"
)

type UserHandlerPort interface {
	GetAll(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
}
