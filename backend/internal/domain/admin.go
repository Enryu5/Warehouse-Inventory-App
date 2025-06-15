package domain

type Admin struct {
	AdminID  int    `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
