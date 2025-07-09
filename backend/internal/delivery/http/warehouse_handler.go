package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
	"github.com/gorilla/mux"
)

type WarehouseHandler struct {
	Usecase usecase.WarehouseUsecase
}

func NewWarehouseHandler(r *mux.Router, uc usecase.WarehouseUsecase) {
	handler := &WarehouseHandler{Usecase: uc}

	r.HandleFunc("/", handler.GetAll).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", handler.GetByID).Methods("GET")
	r.HandleFunc("/", handler.Create).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", handler.Update).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", handler.Delete).Methods("DELETE")
}

func (h *WarehouseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	warehouses, err := h.Usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(warehouses)
}

func (h *WarehouseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	warehouse, err := h.Usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(warehouse)
}

func (h *WarehouseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var warehouse domain.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Usecase.Create(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *WarehouseHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var warehouse domain.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	warehouse.WarehouseID = id
	if err := h.Usecase.Update(&warehouse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WarehouseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.Usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
