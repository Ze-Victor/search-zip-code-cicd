package config

import (
	"fmt"

	"os"

	"gorm.io/gorm"
)

var (
	Secret_key = os.Getenv("Secret_key")
	Base_Path  = os.Getenv("Base_Path")
	Db         *gorm.DB
)

func Init() error {
	var err error

	Db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
