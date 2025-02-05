package domain_contracts_repositories_game_dto

type InputCreateDto struct {
	PlayerOneID      uint
	PlayerTwoID      *uint
	IsAgainstMachine bool
}
