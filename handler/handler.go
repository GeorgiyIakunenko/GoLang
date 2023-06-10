package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"http/data/model"
	"http/data/repository/file"
	"http/response"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	s := file.NewSuppliersRepository()

	jsonData, err := s.GetAll()
	if err != nil {
		fmt.Errorf("Can't get all: %v", err)
		return
	}

	response.SendJson(w, 200, jsonData)

}

func GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplierIDStr, ok := vars["id"]

	if !ok {
		response.SendBadRequestError(w, errors.New("supplier ID not found in path"))
		return
	}

	supplierID, err := strconv.Atoi(supplierIDStr)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	s := file.NewSuppliersRepository()
	supplier, err := s.GetById(supplierID)

	if err != nil {
		http.Error(w, "No supplier with this id", http.StatusBadRequest)
		return
	}

	response.SendJson(w, 200, supplier)
}

func UpdateSupplierById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplierIDStr, ok := vars["id"]
	if !ok {
		response.SendBadRequestError(w, errors.New("supplier ID not found in path"))
		return
	}

	supplierID, err := strconv.Atoi(supplierIDStr)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	defer r.Body.Close()

	var supplier model.Supplier
	err = json.Unmarshal(body, &supplier)

	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	s := file.NewSuppliersRepository()

	err = s.UpdateSupplierById(supplierID, supplier)

	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "SupplierAddedSuccessfully")

}

func CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	var newOder model.Order

	err = json.Unmarshal(body, &newOder)
	if err != nil {
		response.SendBadRequestError(w, err)
		return
	}

	o := file.NewOrdersRepository()
	err = o.Create(newOder)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "OrderAddedSuccessfully")
}

func DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	OrderIdstr, ok := vars["id"]
	if !ok {
		response.SendBadRequestError(w, errors.New("order ID not found in path"))
		return
	}
	o := file.NewOrdersRepository()
	orderId, err := strconv.Atoi(OrderIdstr)
	if err != nil {
		response.SendBadRequestError(w, errors.New("invalid order ID"))
		return
	}

	err = o.DeleteById(orderId)
	if err != nil {
		response.SendServerError(w, err)
		return
	}

	response.SendOK(w, "OrderDeletedSuccessfully")
}

func RefreshSuppliers(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://foodapi.golang.nixdev.co/suppliers")
	if err != nil {
		http.Error(w, fmt.Errorf("Can't fetch data: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		Suppliers []*model.Supplier `json:"suppliers"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		response.SendServerError(w, err)
		return
	}
	s := file.NewSuppliersRepository()
	err = s.RefreshSuppliers(data.Suppliers)
	if err != nil {
		response.SendBadRequestError(w, err)
	}

	response.SendOK(w, "RefreshSuppliers is done")
}
