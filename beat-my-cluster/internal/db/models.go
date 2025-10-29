package db

type Player struct {
	ID      uint     `gorm:"primaryKey"`
	Name    string   `gorm:"unique;not null"`
	Attacks []Attack `gorm:"foreignKey:PlayerID"`
}

type Attack struct {
	ID        uint
	PlayerID  uint `gorm:"not null;index"`
	Intensity int
}
