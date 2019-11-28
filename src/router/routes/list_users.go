package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"user-management/src/types"
)

type ListUsersHandler struct {
	um types.Users
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

func (h ListUsersHandler) isAllowed(u types.User) bool {
	return u.Role() == "admin"
}

func serializeUsers(userList []types.User) ([]byte, error) {
	dtoList := make([]types.UserDTO, 0)

	for _, u := range userList {
		dtoList = append(dtoList, types.UserDTO{
			Username: u.Username(),
			Role:     u.Role(),
		})
	}

	return json.Marshal(dtoList)
}
