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

func (r gameRepository) ListPositions(gameID uint) (*domain_contracts_repositories_game_dto.OutputListPositionsDto, error) {

	var positions []domain_entities.Position

	err := r.db.Table("positions p").
		Joins("INNER JOIN games g ON g.id = p.game_id").
		Where("g.id = ? AND avaible IS true", gameID).
		Order("p.y DESC, p.x ASC").
		Find(&positions).Error

	if err != nil {
		return nil, err
	}

	return &domain_contracts_repositories_game_dto.OutputListPositionsDto{
		Data: func() []domain_contracts_repositories_game_dto.ItemListPositionsDto {
			data := make([]domain_contracts_repositories_game_dto.ItemListPositionsDto, len(positions))

			for i, p := range positions {
				data[i] = domain_contracts_repositories_game_dto.ItemListPositionsDto{
					X:       p.X,
					Y:       p.Y,
					Value:   p.Value,
					Avaible: p.Avaible,
				}
			}
			return data
		}(),
	}, nil
}
func (r gameRepository) NextPlayerToMove(gameID uint) (*domain_entities.Player, error) {
	var game domain_entities.Game

	err := r.db.Model(&domain_entities.Game{}).Where("id = ?", gameID).
		Preload("PlayerOne").Preload("PlayerTwo").
		Take(&game).Error

	if err != nil {
		return nil, err
	}

	if game.ID == 0 {
		return nil, errors.New("game not found")
	}

	var maxMoveOrder *uint

	err = r.db.Table("positions p").
		Select("MAX(p.move_order)").
		Joins("INNER JOIN games g ON g.id = p.game_id").
		Where("p.avaible IS false AND g.id = ?", gameID).
		Scan(&maxMoveOrder).Error

	if err != nil {
		return nil, err
	}

	if maxMoveOrder == nil { // nao houve nenhuma jogada, retorna jogador 1
		return game.PlayerOne, nil
	}

	var lastPlayerID *uint = nil

	err = r.db.Table("positions p").
		Select("p.player_id").
		Joins("INNER JOIN games g ON g.id = p.game_id").
		Where("g.id = ? AND p.avaible IS false AND p.move_order = ?", gameID, *maxMoveOrder).
		Scan(&lastPlayerID).Error

	if err != nil {
		return nil, err
	}

	if lastPlayerID == &game.PlayerOneID {
		return game.PlayerTwo, nil
	}

	return game.PlayerOne, nil
}

func (r gameRepository) Move(gameID uint, playerID uint, position domain_contracts_repositories_game_dto.InputMoveDto) (*domain_contracts_repositories_game_dto.OutputMoveDto, error) {

	var game domain_entities.Game

	err := r.db.Model(&domain_entities.Game{}).
		Where("id = ? AND (player_one_id = ? OR player_two_id = ?)", gameID, playerID, playerID).
		Take(&game).Error

	if err != nil {
		return nil, err
	}

	if game.ID == 0 {
		return nil, errors.New("game not found")
	}

	player, err := r.NextPlayerToMove(game.ID)

	if err != nil {
		return nil, err
	}

	if player.ID != playerID {
		return nil, errors.New("invalid player to move")
	}

	positions, err := r.ListPositions(gameID)
	if err != nil {
		return nil, err
	}

	for _, p := range positions.Data {
		if !p.Avaible && (p.X == position.X && p.Y == position.Y) {
			return nil, errors.New("invalid position to move")
		}
	}

	//fazer logica de verificar se ha ganhador / velha, atualisar o game etc..
	return nil, nil
}
