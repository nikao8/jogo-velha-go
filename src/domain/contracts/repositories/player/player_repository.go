package domain_contracts_repositories_player

import domain_contracts_repositories_player_dto "jogo-velha/src/domain/contracts/repositories/player/dto"

type IPlayerRepository interface {
	Login(inputDto domain_contracts_repositories_player_dto.InputLoginDto) (*domain_contracts_repositories_player_dto.OutputLoginDto, error)
	Create(inputDto domain_contracts_repositories_player_dto.InputCreateDto) error
	Update(playerID uint, inputDto domain_contracts_repositories_player_dto.InputUpdateDto) error
	Games(playerID uint) (*domain_contracts_repositories_player_dto.OutputGamesDto, error)
}
