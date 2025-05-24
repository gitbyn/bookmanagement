package database

import (
	"bookmanagement/config"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.LoadEnv()

	host := config.GetEnv("db_host")
	port := config.GetEnv("db_port")
	user := config.GetEnv("db_user")
	password := config.GetEnv("db_pass")
	dbname := config.GetEnv("db_name")

	// Example: sqlserver://username:password@host:port?database=dbname
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQL Server: ", err)
	}
}