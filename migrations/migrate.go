package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"srbolabApp/database"
)

func MigrateDatabase() {
	database.Connect()
	m, err := migrate.New(
		"file://migrations",
		//fixme fix port for production
		"postgres://postgres:passw0rd@localhost:5433/"+database.DatasourceName+"?sslmode=disable&search_path=public",
	)
	if err != nil {
		log.Fatal(err)
	}
	//FIXME
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}
