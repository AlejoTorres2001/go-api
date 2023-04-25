package main

import (
	"net/http"
	"os"
	"log"
	"github.com/gorilla/mux"
	"github.com/AlejoTorres2001/go-gorm-restapi/internal/routes"
	"github.com/AlejoTorres2001/go-gorm-restapi/internal/db"
	"github.com/joho/godotenv"

)
// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
	}
}

func main() {
	db.DBConnection(db.DbConnectionParams{
		Host: os.Getenv("POSTGRES_DB_HOST"),
		User: os.Getenv("POSTGRES_DB_USER"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
		Dbname: os.Getenv("POSTGRES_DB_NAME"),
		Port: os.Getenv("POSTGRES_DB_PORT"),
		Sslmode: os.Getenv("POSTGRES_DB_SSL"),
		TimeZone: os.Getenv("POSTGRES_DB_TIMEZONE"),
	})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
