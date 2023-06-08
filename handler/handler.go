package handler

import (
	"fmt"
	"http/data/repository/file"
	"net/http"
)

func GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s, err := file.NewSuppliersRepository()
		if err != nil {
			fmt.Errorf("Can't create repository: %v", err)
			return
		}

		jsonData, err := s.GetAll()
		if err != nil {
			fmt.Errorf("Can't get all: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	default:
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
	}
}

/*func ShowSupplier(w http.ResponseWriter, r *http.Request) {
	pattern := regexp.MustCompile("\\d+")
	match := pattern.FindString(r.URL.Path)

	supplierId, err := strconv.Atoi(match)
	if err != nil {
		http.Error(w, fmt.Errorf("%v", err).Error(), http.StatusBadRequest)
		return
	}

}*/
