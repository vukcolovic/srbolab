package main

import (
	"srbolabApp/loger"
	"srbolabApp/migrations"
)

func main() {
	//service.PdfService.CreateCertificate(model.Certificate{})
	loger.InfoLog.Println("Starting application...")
	migrations.MigrateDatabase()
	runServer()
}
