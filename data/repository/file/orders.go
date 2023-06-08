package file

import (
	"encoding/json"
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

func NewOrdersRepository() (*OrdersRepository, error) {
	return &OrdersRepository{
		Directory: "data/orders",
	}, nil
}

func (r *OrdersRepository) Create(order model.Order) error {
	files, err := ioutil.ReadDir(r.Directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return err
	}

	// Find the highest numbered file
	lastFile := files[len(files)-1]
	fileName := lastFile.Name()
	fileNumber, _ := strconv.Atoi(fileName)
	nextNumber := fileNumber + 1

	// Create a new file with the next file number
	filePath := filepath.Join(r.Directory, strconv.Itoa(nextNumber))
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	orderData, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error encoding order:", err)
		return err
	}

	_, err = file.Write(orderData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("File", fileName, "created and content written successfully!")

	return nil
}
