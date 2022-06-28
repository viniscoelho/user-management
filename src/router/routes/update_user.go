package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"user-management/src/types"
	"user-management/src/types/userstore"

	"github.com/gorilla/mux"
)

type updateUser struct {
	um types.Users
}

func NewUpdateUserHandler(um types.Users) *updateUser {
	return &updateUser{um}
}

func (h updateUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

	newUserDTO := types.UserDTO{}
	err = json.Unmarshal(body, &newUserDTO)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	newUser, err := userstore.NewUserFromDTO(newUserDTO)
	if err != nil {
		log.Printf("Error: %s", err)

		switch err.(type) {
		case userstore.InvalidPasswordError:
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
		case userstore.UserDoesNotExistError:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("username does not exist"))
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("internal server error"))
		}

		return
	}
}

func (h updateUser) isAllowed(u types.User, targetUsername string) bool {
	return u.Role() == "admin" || u.Username() == targetUsername
}
