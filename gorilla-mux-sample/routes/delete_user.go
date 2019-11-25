package routes

import (
	"net/http"
	users "user-management/mockgen-sample"
)

type DeleteUserHandler struct {
	um users.Users
}

func (h DeleteUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
