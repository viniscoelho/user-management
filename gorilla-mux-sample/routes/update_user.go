package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	users "user-management/mockgen-sample"

	"github.com/gorilla/mux"
)

type UpdateUserHandler struct {
	um users.Users
}

func (h UpdateUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	newUserDTO := users.UserDTO{}
	err = json.Unmarshal(body, &newUserDTO)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	newUser, err := users.NewUserFromDTO(newUserDTO)
	if err != nil {
		log.Printf("Error: %s", err)

		switch err.(type) {
		case users.InvalidPasswordError:
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("invalid password: it should contain at least 8 characters"))
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("internal server error"))
		}

		return
	}

	err = h.um.UpdateUser(newUser)
	if err != nil {
		log.Printf("Error: %s", err)

		switch err.(type) {
		case users.UserDoesNotExistError:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("username does not exist"))
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("internal server error"))
		}

		return
	}
}

func (h UpdateUserHandler) isAllowed(u users.User, targetUsername string) bool {
	return u.Role() == "admin" || u.Username() == targetUsername
}
