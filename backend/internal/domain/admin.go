package domain

type Admin struct {
	AdminID         int    `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	Username        string `gorm:"not null;unique" json:"username"`
	Hashed_Password string `gorm:"not null" json:"hashed_password"`
}
