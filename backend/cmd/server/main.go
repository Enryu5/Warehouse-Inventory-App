package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/config"
	delivery "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/delivery/http"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/database"
	firebaseUtil "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/firebase"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/persistence"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/usecase"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	db := database.InitPostgres()

	firebaseUtil.InitFirebase()

	// Run migrations only if --migrate flag is passed
	if len(os.Args) > 1 && os.Args[1] == "--migrate" {
		if err := database.Migrate(db); err != nil {
			log.Fatalf("Database migration failed: %v", err)
		}
		return // Exit after migration
	}

	// Initialize repositories
	warehouseRepo := persistence.NewWarehouseRepository(db)
	itemRepo := persistence.NewItemRepository(db)
	adminRepo := persistence.NewAdminRepository(db)
	stockRepo := persistence.NewStockRepository(db)

	// Initialize use cases
	warehouseUC := usecase.NewWarehouseUsecase(warehouseRepo)
	itemUC := usecase.NewItemUsecase(itemRepo)
	adminUC := usecase.NewAdminUsecase(adminRepo)
	stockUC := usecase.NewStockUsecase(stockRepo)

	// Initialize HTTP router with all use cases
	router := delivery.NewRouter(
		warehouseUC,
		itemUC,
		adminUC,
		stockUC,
		db,
	)

	// Add CORS middleware if needed
	handler := addCorsMiddleware(router)

	// Start HTTP server
	serverAddr := ":8080"
	log.Printf("Server is running on %s...\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, handler); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// addCorsMiddleware adds CORS headers to all responses
func addCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
