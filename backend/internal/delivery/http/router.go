package http

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
	"github.com/gorilla/mux"
)

func NewRouter(
	warehouseUC usecase.WarehouseUsecase,
	itemUC usecase.ItemUsecase,
	adminUC usecase.AdminUsecase,
	stockUC usecase.StockUsecase,
) *mux.Router {
	// Create a new router
	router := mux.NewRouter()

	// Create API subrouter with common prefix
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Set up all routes
	setupWarehouseRoutes(apiRouter, warehouseUC)
	setupItemRoutes(apiRouter, itemUC)
	setupAdminRoutes(apiRouter, adminUC)
	setupStockRoutes(apiRouter, stockUC)

	return router
}

func setupWarehouseRoutes(router *mux.Router, warehouseUC usecase.WarehouseUsecase) {
	NewWarehouseHandler(router.PathPrefix("/warehouses").Subrouter(), warehouseUC)
}

func setupItemRoutes(router *mux.Router, itemUC usecase.ItemUsecase) {
	NewItemHandler(router.PathPrefix("/items").Subrouter(), itemUC)
}

func setupAdminRoutes(router *mux.Router, adminUC usecase.AdminUsecase) {
	NewAdminHandler(router.PathPrefix("/admins").Subrouter(), adminUC)
}

func setupStockRoutes(router *mux.Router, stockUC usecase.StockUsecase) {
	NewStockHandler(router.PathPrefix("/stocks").Subrouter(), stockUC)
}
