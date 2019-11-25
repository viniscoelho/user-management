package routes

import (
	"net/http"
	users "user-management/mockgen-sample"
)

type ListUsersHandler struct {
	um users.Users
}

func (h ListUsersHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
