package main

import "srbolabApp/migrations"

func main() {
	migrations.MigrateDatabase()
	runServer()
}
