package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db 		*gorm.DB
	logger 	*Logger
)

func Init() error {
	var err error
	//Initialize SQLite DB
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initialization sqlite: %v", err)
	}
	return nil
}

func GetSQLiteDB() *gorm.DB {
	return	db
}

func GetLogger(p string) *Logger {
	//Init Logger
	logger = NewLogger(p)
	return logger
}