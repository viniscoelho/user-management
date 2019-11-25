package routes

import (
	"net/http"
	users "user-management/mockgen-sample"
)

type UpdateUserHandler struct {
	um users.Users
}

func (h UpdateUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
