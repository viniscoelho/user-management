package main

import (
	"log"
	"net/http"
	"time"

	"user-management/gorilla-mux-sample/routes"
	users "user-management/mockgen-sample"
)

func main() {
	um, err := users.NewUserManagement()
	if err != nil {
		log.Fatal("could not initialize storage")
	}

	s := &http.Server{
		Handler:      routes.CreateRoutes(um),
		ReadTimeout:  0,
		WriteTimeout: 0,
		Addr:         ":3000",
		IdleTimeout:  time.Second * 60,
	}
	log.Fatal(s.ListenAndServe())
}
