package routes

import (
	"encoding/json"
	"log"
	"net/http"
	users "user-management/mockgen-sample"

	"github.com/gorilla/mux"
)

const resourceName = "username"

type ReadUserHandler struct {
	um users.Users
}

func serializeUser(u users.User) ([]byte, error) {
	dto := users.UserDTO{
		Username: u.Username(),
		Role:     u.Role(),
	}
	return json.Marshal(dto)
}

func (h ReadUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != authorizationHeaderToken {
		log.Printf("Unauthorized request to resource: missing authorization header")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("unauthorized"))
		return
	}

	vars := mux.Vars(r)
	username := vars[resourceName]
	u, err := h.um.ReadUser(username)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	content, err := serializeUser(u)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(content)
}
