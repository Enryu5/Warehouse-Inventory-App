package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
	"github.com/gorilla/mux"
)

type ItemHandler struct {
	Usecase usecase.ItemUsecase
}

func NewItemHandler(r *mux.Router, uc usecase.ItemUsecase) {
	handler := &ItemHandler{Usecase: uc}

	r.HandleFunc("/items", handler.GetAll).Methods("GET")
	r.HandleFunc("/items/{id:[0-9]+}", handler.GetByID).Methods("GET")
	r.HandleFunc("/items", handler.Create).Methods("POST")
	r.HandleFunc("/items/{id:[0-9]+}", handler.Update).Methods("PUT")
	r.HandleFunc("/items/{id:[0-9]+}", handler.Delete).Methods("DELETE")
}

func (h *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.Usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	item, err := h.Usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item domain.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Usecase.Create(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var item domain.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.ItemID = id
	if err := h.Usecase.Update(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.Usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
