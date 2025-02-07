package domain_entities

import "time"

type Game struct {
	ID uint `gorm:"primaryKey"`

	WinnerID *uint   `gorm:"null;default:null"`
	Winner   *Player `gorm:"foreignKey:WinnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	PlayerOneID uint    `gorm:"not null"`
	PlayerOne   *Player `gorm:"foreignKey:PlayerOneID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PlayerTwoID uint    `gorm:"not null"`
	PlayerTwo   *Player `gorm:"foreignKey:PlayerTwoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Draw bool `gorm:"not null;default:false"`

	Finished   bool       `gorm:"not null;default:false"`
	FinishedAt *time.Time `gorm:"null;default:null"`

	Started   bool       `gorm:"not null;default:false"`
	StartedAt *time.Time `gorm:"null;default:null"`
}
