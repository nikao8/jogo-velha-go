package domain_contracts_repositories_player_dto

import "time"

type InputLoginDto struct {
	Login    string
	Password string
}

type OutputLoginDto struct {
	ID        uint
	Name      string
	Login     string
	CreatedAt time.Time
}

type InputCreateDto struct {
	Name     string
	Login    string
	Password string
}

type InputUpdateDto struct {
	Name     string
	Login    string
	Password string
}

type ItemGamesDto struct {
	ID            uint
	WinnerName    string
	PlayerOneName string
	PlayerTwoName string
	Draw          bool
	Finished      bool
	FinishedAt    *time.Time
	Started       bool
	StartedAt     *time.Time
}

type OutputGamesDto struct {
	Data []ItemGamesDto
}
