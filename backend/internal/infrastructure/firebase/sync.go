package firebase

import (
	"context"
	"fmt"

	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
)

// --- ITEM SYNC ---

func SyncItemToFirebase(item *domain.Item) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	ref := client.NewRef(fmt.Sprintf("items/%d", item.ItemID))
	return ref.Set(ctx, item)
}

func DeleteItemFromFirebase(itemID int) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	ref := client.NewRef(fmt.Sprintf("items/%d", itemID))
	return ref.Delete(ctx)
}

// --- WAREHOUSE SYNC ---

func SyncWarehouseToFirebase(warehouse *domain.Warehouse) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	ref := client.NewRef(fmt.Sprintf("warehouses/%d", warehouse.WarehouseID))
	return ref.Set(ctx, warehouse)
}

func DeleteWarehouseFromFirebase(warehouseID int) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	ref := client.NewRef(fmt.Sprintf("warehouses/%d", warehouseID))
	return ref.Delete(ctx)
}

// --- STOCK SYNC ---

func SyncStockToFirebase(stock *domain.Stock) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	// Composite key: warehouseID_itemID
	key := fmt.Sprintf("%d_%d", stock.WarehouseID, stock.ItemID)
	ref := client.NewRef(fmt.Sprintf("stocks/%s", key))
	return ref.Set(ctx, stock)
}

func DeleteStockFromFirebase(warehouseID, itemID int) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%d_%d", warehouseID, itemID)
	ref := client.NewRef(fmt.Sprintf("stocks/%s", key))
	return ref.Delete(ctx)
}

// --- ADMIN SYNC ---

func SyncAdminToFirebase(admin *domain.Admin) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	ref := client.NewRef(fmt.Sprintf("admins/%d", admin.AdminID))
	return ref.Set(ctx, admin)
}

func DeleteAdminFromFirebase(adminID int) error {
	ctx := context.Background()
	client, err := App.Database(ctx)
	if err != nil {
		return err
	}
	ref := client.NewRef(fmt.Sprintf("admins/%d", adminID))
	return ref.Delete(ctx)
}
