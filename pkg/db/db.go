package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
}