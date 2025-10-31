package db

type Player struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique;not null"`
	NumAttacks int    `gorm:"default:0"`
}


