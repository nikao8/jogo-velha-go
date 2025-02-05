package main

import (
	"fmt"
	infra_database "jogo-velha/src/infra/database"
)

const RUN_MIGRATE bool = true

func main() {
	fmt.Println("Iniciando Jogo da Velha...")

	db, err := infra_database.GetDBConnection()

	if err != nil {
		panic(err)
	}

	if RUN_MIGRATE {
		if err := infra_database.UpDatabase(db); err != nil {
			panic(err)
		}
	}
}
