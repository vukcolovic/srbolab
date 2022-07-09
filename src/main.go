package main

import (
	"srbolabApp/loger"
	"srbolabApp/migrations"
)

func main() {
	loger.InfoLog.Println("Starting application...")
	migrations.MigrateDatabase()
	runServer()
}
