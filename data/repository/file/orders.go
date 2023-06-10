package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"http/data/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type OrdersRepository struct {
	Directory string `json:"directory"`
}

func NewOrdersRepository() *OrdersRepository {
	return &OrdersRepository{
		Directory: "data/orders",
	}
}

func (r *OrdersRepository) Create(order model.Order) error {
	files, err := ioutil.ReadDir(r.Directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return err
	}

	order.Id = len(files) + 1

	if order.Status != "created" && order.Status != "in_progress" && order.Status != "done" {
		fmt.Println("Invalid status:", order.Status)
		return errors.New("invalid status")
	}

	orderData, err := json.Marshal(order)

	if err != nil {
		fmt.Println("Error encoding order:", err)
		return err
	}

	filePath := filepath.Join(r.Directory, strconv.Itoa(order.Id)+".json")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(orderData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}

func (r *OrdersRepository) DeleteById(id int) error {
	_, err := ioutil.ReadDir(r.Directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return err
	}

	filePath := filepath.Join(r.Directory, strconv.Itoa(id)+".json")
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting from file:", err)
		return err
	}

	return nil
}
