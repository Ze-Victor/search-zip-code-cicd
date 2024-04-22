package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	Secret_key = "secret-key-token"
	Base_Path  = "http://localhost:8001/api/v1"
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
