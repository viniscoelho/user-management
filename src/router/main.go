package main

import (
	"log"
	"net/http"
	"time"
	"user-management/src/router/routes"
	"user-management/src/types/userstore"
)

func main() {
	um, err := userstore.NewUserManagement()
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
