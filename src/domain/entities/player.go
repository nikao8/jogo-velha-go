package domain_entities

import "time"

type Player struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Login     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Machine   bool      `gorm:"not null;default:false"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}
