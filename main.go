package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"http/handler"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/suppliers", handler.GetAllSuppliers).Methods("GET")
	r.HandleFunc(`/suppliers/{id:[0-9]+}`, handler.GetSupplierByID).Methods("GET")
	r.HandleFunc(`/suppliers/{id:[0-9]+}`, handler.UpdateSupplierById).Methods("PUT")
	r.HandleFunc("/orders", handler.CreateNewOrder).Methods("POST")
	r.HandleFunc(`/orders/{id:[0-9]+}`, handler.DeleteOrderById).Methods("DELETE")
	fmt.Printf("server is running")
	http.ListenAndServe(":8080", r)

}

/*resp, err := http.Get("https://foodapi.golang.nixdev.co/suppliers")
if err != nil {
	fmt.Errorf("Can't fetch data: %v", err)
	return
}

body, err := io.ReadAll(resp.Body)
if err != nil {
	panic(err)
}

fmt.Printf(string(body))*/
