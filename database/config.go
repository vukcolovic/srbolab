package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	driverName     = "postgres"
	DatasourceName = "srbolabdb"
)

var (
	Client *sqlx.DB
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		//FIXME port
		"127.0.0.1", 5433, "postgres", "passw0rd", DatasourceName)
	var err error
	Client, err = sqlx.Open(driverName, psqlInfo)
	if err != nil {
		panic(err)
	}
	if errConn := Client.Ping(); errConn != nil {
		panic(errConn)
	}
	//FIXME log properly
	log.Println("Database succesfully connected")
}
