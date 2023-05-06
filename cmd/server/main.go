package main

import (
	"fmt"
	"net/http"

	"github.com/buraksenn/expense-tracker/internal/service"
)

func main() {
	service := service.New()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: service.Router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
