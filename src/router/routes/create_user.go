package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"user-management/src/types"
	"user-management/src/types/userstore"
)

type CreateUserHandler struct {
	um types.Users
}

func (h CreateUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	requesterName := r.Header.Get("Authorization")
	if len(requesterName) == 0 {
		log.Printf("Unauthorized request to resource: missing authorization header")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("unauthorized"))
		return
	}

	requester, err := h.um.ReadUser(requesterName)
	if err != nil || !h.isAllowed(requester) {
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
		case userstore.InvalidPasswordError, userstore.InvalidUsernameError, userstore.InvalidRoleError:
			rw.WriteHeader(http.StatusBadRequest)
			message := "invalid fields: username should have at least one character, " +
				"password should contain at least 8 characters and role must be either admin or regular"
			rw.Write([]byte(message))
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("internal server error"))
		}

		return
	}

	err = h.um.CreateUser(newUser)
	if err != nil {
		log.Printf("Error: %s", err)

		switch err.(type) {
		case userstore.UserAlreadyExistsError:
			rw.WriteHeader(http.StatusConflict)
			rw.Write([]byte("username already taken"))
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("internal server error"))
		}

		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (h CreateUserHandler) isAllowed(u types.User) bool {
	return u.Role() == "admin"
}
