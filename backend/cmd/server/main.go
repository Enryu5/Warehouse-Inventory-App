package main

import (
	"log"
	"net/http"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/config"
	delivery "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/delivery/http"
)

func main() {
	config.LoadEnv()
	db := config.ConnectPostgres()
	defer db.Close()

	mux := http.NewServeMux()
	delivery.RegisterRoutes(mux)

	log.Println("✅ Server is running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
