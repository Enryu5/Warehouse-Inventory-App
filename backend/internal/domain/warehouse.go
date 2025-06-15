package domain

type Warehouse struct {
	WarehouseID   int    `gorm:"primaryKey;autoIncrement" json:"warehouse_id"`
	AdminID       int    `gorm:"not null" json:"admin_id"`
	WarehouseName string `gorm:"not null" json:"warehouse_name"`
	Location      string `gorm:"not null" json:"location"`
}
