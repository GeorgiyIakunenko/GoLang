package configs

import (
	"auth/internal/server"
	"log"
	"net/http"
)

func Start() {

	cfg := NewConfig()

	authHandler := server.NewAuthHandler(cfg)
	userHandler := server.NewUserHandler(cfg)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
