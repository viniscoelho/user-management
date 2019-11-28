package routes

import (
	"encoding/json"
	"log"
	"net/http"
	users "user-management/mockgen-sample"
)

type ListUsersHandler struct {
	um users.Users
}

func (h ListUsersHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

	u, err := h.um.ListUsers()
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	content, err := serializeUsers(u)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal server error"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(content)
}

func (h ListUsersHandler) isAllowed(u users.User) bool {
	return u.Role() == "admin"
}

func serializeUsers(userList []users.User) ([]byte, error) {
	dtoList := make([]users.UserDTO, 0)

	for _, u := range userList {
		dtoList = append(dtoList, users.UserDTO{
			Username: u.Username(),
			Role:     u.Role(),
		})
	}

	return json.Marshal(dtoList)
}
