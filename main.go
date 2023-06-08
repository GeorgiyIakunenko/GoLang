package main

import (
	"fmt"
	"http/handler"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/suppliers", handler.GetAllSuppliers)

	mux.HandleFunc("/suppliers/2", handler.GetAllSuppliers)

	resp, err := http.Get("https://foodapi.golang.nixdev.co/suppliers")
	if err != nil {
		fmt.Errorf("Can't fetch data: %v", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(body))

	http.ListenAndServe(":8080", mux)

	fmt.Printf("server is running")
}
