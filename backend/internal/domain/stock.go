package domain

import "time"

type Stock struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	WarehouseID int       `gorm:"not null" json:"warehouse_id"`
	ItemID      int       `gorm:"not null" json:"item_id"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	LastUpdated time.Time `gorm:"autoUpdateTime" json:"last_updated"`
}
