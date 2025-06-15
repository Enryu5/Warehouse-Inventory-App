package domain

type Item struct {
	ItemID   int    `gorm:"primaryKey;autoIncrement" json:"item_id"`
	ItemName string `gorm:"not null" json:"item_name"`
}
