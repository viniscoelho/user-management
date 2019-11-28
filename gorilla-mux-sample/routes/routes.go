package routes

import (
	"net/http"

	users "user-management/mockgen-sample"

	"github.com/gorilla/mux"
)

const usernameRouteVar = "username"

func CreateRoutes(um users.Users) *mux.Router {
	r := mux.NewRouter()

	r.Path("/users").
		Methods(http.MethodGet).
		Name("ListUsers").
		Handler(ListUsersHandler{um})
	r.Path("/users").
		Methods(http.MethodPost).
		Name("CreateUser").
		Handler(CreateUserHandler{um})
	r.Path("/users/{username}").
		Methods(http.MethodGet).
		Name("ReadUser").
		Handler(ReadUserHandler{um})
	r.Path("/users/{username}").
		Methods(http.MethodPatch).
		Name("UpdateUser").
		Handler(UpdateUserHandler{um})
	r.Path("/users/{username}").
		Methods(http.MethodDelete).
		Name("DeleteUser").
		Handler(DeleteUserHandler{um})

	return r
}
