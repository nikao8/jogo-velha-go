package domain_entities

type Symbol int

const (
	X Symbol = -1
	O Symbol = 1
)

// Positions criadas quando inicia o jogo (3x3)
type Position struct {
	X uint `gorm:"primaryKey"`
	Y uint `gorm:"primaryKey"`

	GameID uint `gorm:"primaryKey"`
	Game   Game `gorm:"foreignKey:GameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Value    *Symbol `gorm:"null;default:null"`
	PlayerID *uint   `gorm:"null"`
	Player   *Player `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Avaible   bool `gorm:"not null;default:true"`
	MoveOrder uint `gorm:"not null;default:0"`
}
