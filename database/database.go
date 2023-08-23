package database

import (
	"fmt"
	"os"

	"gorm-preload/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Connecting the database and creating a singleton
of the instance so we can use it everywhere
*/
var (
	// DBConn is a database connection singleton
	DBConn *gorm.DB
)

func InitDb() *gorm.DB {
	dsn, envExists := os.LookupEnv("DB_DSN")
	if !envExists {
		panic("Could not connect to the database")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not initialize database")
	}
	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Post{}, &models.Reply{})
	if err != nil {
		fmt.Print("Could not init database")
		return
	}
	fmt.Print("Database migrated")
}
