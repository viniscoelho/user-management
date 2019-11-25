package routes

import (
	"net/http"
	users "user-management/mockgen-sample"
)

type ReadUserHandler struct {
	um users.Users
}

func (h ReadUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
