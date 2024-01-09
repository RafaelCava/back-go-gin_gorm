package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/RafaelCava/kitkit-back-go/cmd/domain/models"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {
	//we read our .env file
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT")) // don't forget to convert int since port is int type.
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Fortaleza",
		host, user, pass, dbname, port)
	db, errSql := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ")
	} else {
		Db = db
		sqlDB, _ := Db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		db.AutoMigrate(&models.User{})
		fmt.Println("Successfully connected to database!")
	}
}
