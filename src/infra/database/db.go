package infra_database

import (
	domain_entities "jogo-velha/src/domain/entities"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=830083 dbname=jogo_velha port=5432 TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // Altere a saída de log conforme necessário
			logger.Config{
				SlowThreshold:             time.Second, // Limite para consultas lentas
				LogLevel:                  logger.Info, // Nível de log (Info exibirá SQLs)
				IgnoreRecordNotFoundError: true,        // Ignorar erros de registro não encontrado
				Colorful:                  true,        // Ativar cores no log
			},
		),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func UpDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain_entities.Player{},
		&domain_entities.Game{},
		&domain_entities.Move{},
		&domain_entities.Position{},
	)

	if err != nil {
		return err
	}

	var notExistsMachine bool

	err = db.Raw(`
	SELECT NOT EXISTS(
		SELECT 1 
		FROM players
		WHERE machine IS true
		LIMIT 1
	)
	`).Scan(&notExistsMachine).Error

	if err != nil {
		return err
	}

	if notExistsMachine {
		err = db.Create(&domain_entities.Player{
			Name:      "Máquina",
			Login:     "machine",
			Password:  "machine",
			Machine:   true,
			CreatedAt: time.Now(),
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}
