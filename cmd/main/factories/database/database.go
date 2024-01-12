// cmd/main/factories/database/database.go - Camada Principal - Factories

package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseOpenConnection() (*gorm.DB, error) {
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
