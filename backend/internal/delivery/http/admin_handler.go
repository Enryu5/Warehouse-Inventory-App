package http

import (
	"encoding/json"
	"net/http"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler struct {
	Usecase usecase.AdminUsecase
}

func NewAdminHandler(r *mux.Router, uc usecase.AdminUsecase) {
	handler := &AdminHandler{Usecase: uc}

	r.HandleFunc("/{username}", handler.GetByUsername).Methods("GET")
	r.HandleFunc("/", handler.Create).Methods("POST")
}

func (h *AdminHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	admin, err := h.Usecase.GetByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func (h *AdminHandler) Create(w http.ResponseWriter, r *http.Request) {
	var admin domain.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before saving
	hashed, err := bcrypt.GenerateFromPassword([]byte(admin.Hashed_Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	admin.Hashed_Password = string(hashed)

	if err := h.Usecase.Create(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
