package routes

import (
	"log"
	"net/http"
	"user-management/src/types"
	"user-management/src/types/userstore"

	"github.com/gorilla/mux"
)

type DeleteUserHandler struct {
	um types.Users
}

func (h DeleteUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

	err = h.um.DeleteUser(targetUsername)
	if err != nil {
		log.Printf("Error: %s", err)

		switch err.(type) {
		case userstore.UserDoesNotExistError:
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("user does not exist"))
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("internal server error"))
		}

		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (h DeleteUserHandler) isAllowed(u types.User, targetUsername string) bool {
	return u.Username() != targetUsername && u.Role() == "admin"
}
