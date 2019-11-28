package routes

import (
	"encoding/json"
	"log"
	"net/http"
	users "user-management/mockgen-sample"

	"github.com/gorilla/mux"
)

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
	requesterName := r.Header.Get("Authorization")
	if len(requesterName) == 0 {
		log.Printf("Unauthorized request to resource: missing authorization header")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("unauthorized"))
		return
	}

	vars := mux.Vars(r)
	targetUsername := vars[usernameRouteVar]

	requester, err := h.um.ReadUser(requesterName)
	if err != nil || !h.isAllowed(requester, targetUsername) {
		log.Printf("Insufficient authorization for this operation")
		rw.WriteHeader(http.StatusForbidden)
		rw.Write([]byte("forbidden"))
		return
	}

	u, err := h.um.ReadUser(targetUsername)
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

func (h ReadUserHandler) isAllowed(u users.User, targetUsername string) bool {
	return u.Role() == "admin" || u.Username() == targetUsername
}
