package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"user-management/src/types"

	"github.com/gorilla/mux"
)

type readUser struct {
	um types.Users
}

func NewReadUserHandler(um types.Users) *readUser {
	return &readUser{um}
}

func serializeUser(u types.User) ([]byte, error) {
	dto := types.UserDTO{
		Username: u.Username(),
		Role:     u.Role(),
	}
	return json.Marshal(dto)
}

func (h readUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

func (h readUser) isAllowed(u types.User, targetUsername string) bool {
	return u.Role() == "admin" || u.Username() == targetUsername
}
