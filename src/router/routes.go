package router

import (
	"net/http"

	"user-management/src/router/routes"
	"user-management/src/types"

	"github.com/gorilla/mux"
)

func CreateRoutes(um types.Users) *mux.Router {
	r := mux.NewRouter()

	r.Path("/users").
		Methods(http.MethodGet).
		Name("ListUsers").
		Handler(routes.NewListUsersHandler(um))
	r.Path("/users").
		Methods(http.MethodPost).
		Name("CreateUser").
		Handler(routes.NewCreateUserHandler(um))
	r.Path("/users/{username}").
		Methods(http.MethodGet).
		Name("ReadUser").
		Handler(routes.NewReadUserHandler(um))
	r.Path("/users/{username}").
		Methods(http.MethodPatch).
		Name("UpdateUser").
		Handler(routes.NewUpdateUserHandler(um))
	r.Path("/users/{username}").
		Methods(http.MethodDelete).
		Name("DeleteUser").
		Handler(routes.NewDeleteUserHandler(um))

	return r
}
