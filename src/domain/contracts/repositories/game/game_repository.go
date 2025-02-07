package domain_contracts_repositories_game

import (
	domain_contracts_repositories_game_dto "jogo-velha/src/domain/contracts/repositories/game/dto"
	domain_entities "jogo-velha/src/domain/entities"
)

type IGameRepository interface {
	Create(inputDto domain_contracts_repositories_game_dto.InputCreateDto) (gameID uint, err error)
	Start(gameID uint) error
	ListPositions(gameID uint) (*domain_contracts_repositories_game_dto.OutputListPositionsDto, error)
	NextPlayerToMove(gameID uint) (*domain_entities.Player, error)
	Move(gameID uint, playerID uint, position domain_contracts_repositories_game_dto.InputMoveDto) (*domain_contracts_repositories_game_dto.OutputMoveDto, error)
}
