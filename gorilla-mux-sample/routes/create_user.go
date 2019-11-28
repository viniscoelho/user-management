package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	users "user-management/mockgen-sample"
)

type CreateUserHandler struct {
	um users.Users
}

func (h CreateUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != authorizationHeaderToken {
		log.Printf("Unauthorized request to resource: missing authorization header")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("unauthorized"))
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

	newUser := users.NewUserFromDTO(newUserDTO)
	err = h.um.CreateUser(newUser)
	if err != nil {
		log.Printf("Error: %s", err)

		switch err.(type) {
		case users.UserAlreadyExistsError:
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
