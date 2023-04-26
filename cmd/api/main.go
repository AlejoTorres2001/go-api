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

func init() {
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
	db.InitDatabase(db.DB)


	r := mux.NewRouter()
	routes.RegisterHandlers(routes.UsersHandlers,routes.UsersSubPath,routes.UsersIdentifierPath,r)
	http.ListenAndServe(":3000", r)
}
