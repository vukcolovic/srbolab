package main

import (
	"srbolabApp/loger"
	"srbolabApp/migrations"
)

func main() {
	loger.InfoLog.Println("Starting application...")
	migrations.MigrateDatabase()

	//cert, _ := service.CertificateService.GetCertificateById(2)
	//service.PdfService.CreateCertificate(cert, "")
	runServer()
}
