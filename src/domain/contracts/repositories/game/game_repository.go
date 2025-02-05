package domain_contracts_repositories_game

import domain_contracts_repositories_game_dto "jogo-velha/src/domain/contracts/repositories/game/dto"

type IGameRepository interface {
	Create(inputDto domain_contracts_repositories_game_dto.InputCreateDto) (gameID uint, err error)
	Start(gameID uint) error
}
