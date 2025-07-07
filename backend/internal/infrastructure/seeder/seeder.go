package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/config"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/database"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	config.LoadEnv()
	db := database.InitPostgres()

	reset := len(os.Args) > 1 && os.Args[1] == "--reset"
	if reset {
		log.Println("Resetting all tables...")
		db.Exec("DELETE FROM stocks")
		db.Exec("DELETE FROM warehouses")
		db.Exec("DELETE FROM items")
		db.Exec("DELETE FROM admins")
		log.Println("All tables reset!")
	}

	// Seed Admins
	admins := []domain.Admin{}
	for i := 1; i <= 3; i++ {
		password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := domain.Admin{
			Username:        "admin" + string('0'+i),
			Hashed_Password: string(password),
		}
		db.FirstOrCreate(&admin, domain.Admin{Username: admin.Username})
		admins = append(admins, admin)
	}

	// Seed Warehouses (assign 2 to admin1, 2 to admin2, 1 to admin3)
	warehouses := []domain.Warehouse{
		{WarehouseName: "Warehouse A", Location: "New York", AdminID: admins[0].AdminID},
		{WarehouseName: "Warehouse B", Location: "San Francisco", AdminID: admins[0].AdminID},
		{WarehouseName: "Warehouse C", Location: "Chicago", AdminID: admins[1].AdminID},
		{WarehouseName: "Warehouse D", Location: "Los Angeles", AdminID: admins[1].AdminID},
		{WarehouseName: "Warehouse E", Location: "Houston", AdminID: admins[2].AdminID},
	}
	for _, w := range warehouses {
		db.FirstOrCreate(&w, domain.Warehouse{WarehouseName: w.WarehouseName})
	}

	// Seed Items
	itemNames := []string{"Item A", "Item B", "Item C", "Item D", "Item E", "Item F", "Item G", "Item H", "Item I", "Item J"}
	items := []domain.Item{}
	for _, name := range itemNames {
		it := domain.Item{ItemName: name}
		db.FirstOrCreate(&it, domain.Item{ItemName: name})
		items = append(items, it)
	}

	// Get updated warehouse and item IDs
	var allWarehouses []domain.Warehouse
	var allItems []domain.Item
	db.Find(&allWarehouses)
	db.Find(&allItems)

	// Seed Stocks - 20 random combinations
	rand.Seed(time.Now().UnixNano())
	usedComb := map[string]bool{}
	stockCount := 0
	for stockCount < 20 {
		w := allWarehouses[rand.Intn(len(allWarehouses))]
		it := allItems[rand.Intn(len(allItems))]
		key := fmt.Sprintf("%d-%d", w.WarehouseID, it.ItemID)
		if usedComb[key] {
			continue // skip duplicates
		}
		usedComb[key] = true
		stock := domain.Stock{
			WarehouseID: w.WarehouseID,
			ItemID:      it.ItemID,
			Quantity:    rand.Intn(100) + 1,
			LastUpdated: time.Now(),
		}
		db.Create(&stock)
		stockCount++
	}

	log.Println("Database seeded with 3 admins, 5 warehouses, 10 items, and 20 random stocks.")
}
