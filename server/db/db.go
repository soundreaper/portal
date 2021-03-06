package db

import (
	"fmt"
	"log"

	"github.com/soundreaper/portal/config"
	"github.com/soundreaper/portal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	// get db config from environment
	dbConfig := config.GetDBConfig()
	// connect to database
	database := getConnection(dbConfig)

	// migrate our models
	err := database.AutoMigrate(models.User{}, models.Image{})
	if err != nil {
		log.Fatal("db: error migrating models. err: ", err)
	}

	return database
}

// getConnection will just get the connection to the db with the given config
func getConnection(c *config.DBConfiguration) *gorm.DB {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", c.Username, c.Password, c.DBName, c.Port, c.Host)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("DB Connection Error: %v", err)
	}

	return db
}
