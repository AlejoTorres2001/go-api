package db

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnectionParams struct  
{
	Host string
	User string
	Password string
	Dbname string
	Port string
	Sslmode string
	TimeZone string
}

var DB *gorm.DB

func DBConnection (params DbConnectionParams) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", params.Host, params.User, params.Password, params.Dbname, params.Port, params.Sslmode, params.TimeZone)
	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("Successfully connected to database!")
	}
}
