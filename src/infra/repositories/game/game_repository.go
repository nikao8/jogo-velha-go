package infra_repositories_game

import (
	"errors"
	domain_contracts_repositories_game_dto "jogo-velha/src/domain/contracts/repositories/game/dto"
	domain_entities "jogo-velha/src/domain/entities"
	"time"

	"gorm.io/gorm"
)

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *gameRepository {
	return &gameRepository{db: db}
}

func (r gameRepository) Create(inputDto domain_contracts_repositories_game_dto.InputCreateDto) (gameID uint, err error) {

	var playerOne domain_entities.Player
	var playerTwo domain_entities.Player

	err = r.db.Model(&domain_entities.Player{}).
		Where("id = ?", inputDto.PlayerOneID).
		Take(&playerOne).Error

	if err != nil {
		return
	}

	if playerOne.ID == 0 {
		return 0, errors.New("player one not found")
	}

	if inputDto.IsAgainstMachine {
		err = r.db.Model(&domain_entities.Player{}).
			Where("machine IS TRUE").
			Take(&playerTwo).Error

		if err != nil {
			return
		}

		if playerTwo.ID == 0 {
			return 0, errors.New("machine player not found")
		}

	} else {

		err = r.db.Model(&domain_entities.Player{}).
			Where("id = ?", *inputDto.PlayerTwoID).
			Take(&playerTwo).Error

		if err != nil {
			return
		}

		if playerTwo.ID == 0 {
			return 0, errors.New("player two not found")
		}

	}

	game := domain_entities.Game{
		WinnerID:    nil,
		PlayerOneID: playerOne.ID,
		PlayerTwoID: playerTwo.ID,
		Draw:        false,
		Finished:    false,
		Started:     false,
		StartedAt:   nil,
	}

	err = r.db.Create(&game).Error

	gameID = game.ID
	return
}

func (r gameRepository) Start(gameID uint) error {
	var game domain_entities.Game

	err := r.db.Model(&domain_entities.Game{}).
		Where("id = ?", gameID).
		Take(&game).Error

	if err != nil {
		return err
	}

	if game.ID == 0 {
		return errors.New("game not found")
	}

	tx := r.db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	var positionsToCreate []domain_entities.Position

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			positionsToCreate = append(positionsToCreate, domain_entities.Position{
				X:      uint(i),
				Y:      uint(j),
				GameID: game.ID,
				Value:  nil,
			})
		}
	}

	err = tx.Create(&positionsToCreate).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	now := time.Now()
	game.Started = true
	game.StartedAt = &now

	err = tx.Updates(&game).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
