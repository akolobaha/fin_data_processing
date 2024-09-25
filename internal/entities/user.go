package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Telegram string `gorm:"not null"`
}
