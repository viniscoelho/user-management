package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Path("/users").
		Methods(http.MethodGet).
		Name("ListUsers").
		Handler(ListUsersHandler{})
	r.Path("/users").
		Methods(http.MethodPost).
		Name("CreateUser").
		Handler(CreateUserHandler{})
	r.Path("/users/{username}").
		Methods(http.MethodGet).
		Name("ReadUser").
		Handler(ReadUserHandler{})
	r.Path("/users/{username}").
		Methods(http.MethodPatch).
		Name("UpdateUser").
		Handler(UpdateUserHandler{})
	r.Path("/users/{username}").
		Methods(http.MethodDelete).
		Name("DeleteUser").
		Handler(DeleteUserHandler{})

	return r
}
