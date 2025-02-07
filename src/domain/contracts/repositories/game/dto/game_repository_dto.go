package domain_contracts_repositories_game_dto

import (
	domain_entities "jogo-velha/src/domain/entities"
)

type InputCreateDto struct {
	PlayerOneID      uint
	PlayerTwoID      *uint
	IsAgainstMachine bool
}

type ItemListPositionsDto struct {
	X       uint
	Y       uint
	Value   *domain_entities.Symbol
	Avaible bool
}

type OutputListPositionsDto struct {
	Data []ItemListPositionsDto
}

type OutputMoveDto struct {
	Winner   *domain_entities.Player
	Finished bool
	Draw     *bool
}

type InputMoveDto struct {
	X      uint
	Y      uint
	Symbol domain_entities.Symbol
}
