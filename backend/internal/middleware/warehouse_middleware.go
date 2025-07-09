package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type WarehouseMiddleware struct {
	DB *gorm.DB
}

func NewWarehouseMiddleware(db *gorm.DB) *WarehouseMiddleware {
	return &WarehouseMiddleware{DB: db}
}

// getWarehouseIDFromRequest extracts warehouse_id from URL params or request body
func (wm *WarehouseMiddleware) getWarehouseIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	if warehouseID, exists := vars["id"]; exists {
		return strconv.Atoi(warehouseID)
	}
	if warehouseID, exists := vars["warehouse_id"]; exists {
		return strconv.Atoi(warehouseID)
	}
	return 0, fmt.Errorf("warehouse_id not found in URL")
}

// WarehouseWriteAuthMiddleware checks if the admin has write access to the warehouse
func (wm *WarehouseMiddleware) WarehouseWriteAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth check for GET requests (read-only)
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		// Get admin ID from context (set by JWT middleware)
		adminID, ok := r.Context().Value(AdminIDKey).(int)
		if !ok {
			http.Error(w, "Unauthorized: missing admin ID", http.StatusUnauthorized)
			return
		}

		// Get warehouse ID from request
		warehouseID, err := wm.getWarehouseIDFromRequest(r)
		if err != nil {
			http.Error(w, "Invalid warehouse ID", http.StatusBadRequest)
			return
		}

		// Check if warehouse exists
		var warehouse domain.Warehouse
		result := wm.DB.First(&warehouse, warehouseID)
		if result.Error != nil {
			http.Error(w, "Warehouse not found", http.StatusNotFound)
			return
		}

		// Check if admin is authorized for this warehouse
		if warehouse.AdminID != adminID {
			http.Error(w, "Forbidden: you don't have write access to this warehouse", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
