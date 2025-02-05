package domain_entities

type Symbol int

const (
	X Symbol = -1
	O Symbol = 1
)

type Move struct {
	ID uint `gorm:"primaryKey"`

	GameID uint `gorm:"not null"`
	Game   Game `gorm:"foreignKey:GameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	PositionID uint     `gorm:"not null"`
	Position   Position `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Symbol Symbol `gorm:"not null"`

	PlayerID uint   `gorm:"not null"`
	Player   Player `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
