package routes

import (
	"net/http"
	users "user-management/mockgen-sample"
)

type CreateUserHandler struct {
	um users.Users
}

func (h CreateUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
