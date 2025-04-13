package configs

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("Error initializing postgres: %v", err)
	}

	return nil

}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
