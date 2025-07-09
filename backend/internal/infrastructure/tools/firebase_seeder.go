package main

import (
	"log"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/config"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/database"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/firebase"
)

func main() {
	config.LoadEnv()
	db := database.InitPostgres()
	firebase.InitFirebase()

	// Sync Admins
	var admins []domain.Admin
	db.Find(&admins)
	for _, admin := range admins {
		if err := firebase.SyncAdminToFirebase(&admin); err != nil {
			log.Printf("Failed to sync admin %d: %v", admin.AdminID, err)
		}
	}

	// Sync Warehouses
	var warehouses []domain.Warehouse
	db.Find(&warehouses)
	for _, warehouse := range warehouses {
		if err := firebase.SyncWarehouseToFirebase(&warehouse); err != nil {
			log.Printf("Failed to sync warehouse %d: %v", warehouse.WarehouseID, err)
		}
	}

	// Sync Items
	var items []domain.Item
	db.Find(&items)
	for _, item := range items {
		if err := firebase.SyncItemToFirebase(&item); err != nil {
			log.Printf("Failed to sync item %d: %v", item.ItemID, err)
		}
	}

	// Sync Stocks
	var stocks []domain.Stock
	db.Find(&stocks)
	for _, stock := range stocks {
		if err := firebase.SyncStockToFirebase(&stock); err != nil {
			log.Printf("Failed to sync stock ID %d: %v", stock.ID, err)
		}
	}

	log.Println("Full sync from Postgres to Firebase finished!")
}
