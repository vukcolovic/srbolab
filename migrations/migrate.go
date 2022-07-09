package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"srbolabApp/database"
	"srbolabApp/loger"
)

func MigrateDatabase() {
	database.Connect()
	m, err := migrate.New(
		"file://migrations",
		//fixme fix port for production
		"postgres://postgres:passw0rd@localhost:5433/"+database.DatasourceName+"?sslmode=disable&search_path=public",
	)
	if err != nil {
		loger.ErrorLog.Println("Migrate database error: ", err)
	}

	if err = m.Up(); err != nil && err.Error() != "no change" {
		loger.ErrorLog.Println("Migrate database error: ", err)
	}
}
