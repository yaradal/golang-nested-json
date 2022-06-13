package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/yaradal/flaconi-challenge/nesting"
)

// We hardcode the user/password because not specified in the requirements.
const (
	username = "yara"
	password = "555"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	nestingSvc := nesting.NewService()
	nestingCtrl := nesting.NewController(nestingSvc)

	router := chi.NewRouter()
	router.Use(middleware.BasicAuth("", map[string]string{username: password}))
	router.Mount("/nesting", nestingCtrl)

	fmt.Println("Listening on :8080")

	return http.ListenAndServe(":8080", router)
}
