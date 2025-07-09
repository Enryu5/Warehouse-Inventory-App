package http

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/middleware"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(
	warehouseUC usecase.WarehouseUsecase,
	itemUC usecase.ItemUsecase,
	adminUC usecase.AdminUsecase,
	stockUC usecase.StockUsecase,
	db *gorm.DB,
) *mux.Router {
	// Create a new router
	router := mux.NewRouter()

	// Create API subrouter with common prefix
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Set up login route before applying JWT middleware (so /login is unprotected)
	setupAuthRoutes(apiRouter, adminUC)

	// Apply JWT middleware to all other /api routes
	apiRouter.Use(middleware.JWTMiddleware)

	// Initialize middlewares
	warehouseMW := middleware.NewWarehouseMiddleware(db)

	// Set up all routes under /api
	setupWarehouseRoutes(apiRouter, warehouseUC, warehouseMW)
	setupItemRoutes(apiRouter, itemUC)
	setupAdminRoutes(apiRouter, adminUC)
	setupStockRoutes(apiRouter, stockUC)

	return router
}

func setupAuthRoutes(router *mux.Router, adminUC usecase.AdminUsecase) {
	authHandler := &AuthHandler{AdminUsecase: adminUC}
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
}

func setupWarehouseRoutes(router *mux.Router, warehouseUC usecase.WarehouseUsecase, warehouseMW *middleware.WarehouseMiddleware) {
	// Create a subrouter for warehouse routes
	warehouseRouter := router.PathPrefix("/warehouses").Subrouter()

	// Apply warehouse authorization middleware for write operations
	warehouseRouter.Use(warehouseMW.WarehouseWriteAuthMiddleware)

	// Set up warehouse routes
	NewWarehouseHandler(warehouseRouter, warehouseUC)
}

func setupItemRoutes(router *mux.Router, itemUC usecase.ItemUsecase) {
	NewItemHandler(router.PathPrefix("/items").Subrouter(), itemUC)
}

func setupAdminRoutes(router *mux.Router, adminUC usecase.AdminUsecase) {
	adminRouter := router.PathPrefix("/admins").Subrouter()
	NewAdminHandler(adminRouter, adminUC)
}

func setupStockRoutes(router *mux.Router, stockUC usecase.StockUsecase) {
	NewStockHandler(router.PathPrefix("/stocks").Subrouter(), stockUC)
}
