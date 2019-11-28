package routes

import (
	"log"
	"net/http"
	users "user-management/mockgen-sample"

	"github.com/gorilla/mux"
)

type DeleteUserHandler struct {
	um users.Users
}

func (h DeleteUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != authorizationHeaderToken {
		log.Printf("Unauthorized request to resource: missing authorization header")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("unauthorized"))
		return
	}

	vars := mux.Vars(r)
	username := vars[resourceName]
	err := h.um.DeleteUser(username)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
