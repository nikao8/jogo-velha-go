package domain_entities

// Positions criadas quando inicia o jogo (3x3)
type Position struct {
	X uint `gorm:"primaryKey"`
	Y uint `gorm:"primaryKey"`

	GameID uint `gorm:"primaryKey"`
	Game   Game `gorm:"foreignKey:GameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Value *Symbol `gorm:"null;default:null"`
}
