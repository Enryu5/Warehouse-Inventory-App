package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
	"github.com/gorilla/mux"
)

type StockHandler struct {
	Usecase usecase.StockUsecase
}

func NewStockHandler(r *mux.Router, uc usecase.StockUsecase) {
	handler := &StockHandler{Usecase: uc}

	r.HandleFunc("/warehouse/{warehouse_id:[0-9]+}", handler.GetByWarehouse).Methods("GET")
	r.HandleFunc("/item/{item_id:[0-9]+}", handler.GetByItem).Methods("GET")
	r.HandleFunc("/", handler.Upsert).Methods("POST", "PUT")
	r.HandleFunc("/item/{item_id:[0-9]+}/warehouse/{warehouse_id:[0-9]+}", handler.DeleteByItemAndWarehouse).Methods("DELETE")
}

func (h *StockHandler) GetByWarehouse(w http.ResponseWriter, r *http.Request) {
	warehouseID, _ := strconv.Atoi(mux.Vars(r)["warehouse_id"])
	stocks, err := h.Usecase.GetByWarehouse(warehouseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stocks)
}

func (h *StockHandler) GetByItem(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["item_id"])
	stocks, err := h.Usecase.GetByItem(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stocks)
}

func (h *StockHandler) Upsert(w http.ResponseWriter, r *http.Request) {
	var stock domain.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Usecase.Upsert(&stock); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *StockHandler) DeleteByItemAndWarehouse(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["item_id"])
	warehouseID, _ := strconv.Atoi(mux.Vars(r)["warehouse_id"])
	if err := h.Usecase.DeleteByItemAndWarehouse(itemID, warehouseID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
