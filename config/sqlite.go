package config

import (
	"os"

	"github.com/renatocardoso243/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	//Check if database exists
	dbPatch := "./db/main.db" //Variable to store database path
	_, err := os.Stat(dbPatch)
	if os.IsNotExist(err) {
		logger.Info("Database not found. Creating...")
		//Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPatch)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	
	//Create database and connect
	db, err:= gorm.Open(sqlite.Open(dbPatch), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//Migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Database migration error: %v", err)
		return nil, err
	}
	//Return database
	return db, nil
}