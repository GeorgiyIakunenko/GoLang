package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"http/handler"
	"http/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// use gorilla mux
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/suppliers", handler.GetAllSuppliers).Methods("GET")
	r.HandleFunc(`/suppliers/{id:[0-9]+}`, handler.GetSupplierByID).Methods("GET")
	r.HandleFunc("/suppliers/refresh", handler.RefreshSuppliers).Methods("GET")
	r.HandleFunc(`/suppliers/{id:[0-9]+}`, handler.UpdateSupplierById).Methods("PUT")
	r.HandleFunc("/orders", handler.CreateNewOrder).Methods("POST")
	r.HandleFunc(`/orders/{id:[0-9]+}`, handler.DeleteOrderById).Methods("DELETE")
	r.Use(middleware.LogMiddleware)
	fmt.Printf("server is running")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		err := srv.ListenAndServeTLS("certs/localhost.pem", "certs/localhost-key.pem")
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("shutting down")
	os.Exit(0)

}
