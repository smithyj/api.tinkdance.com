package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(config Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
}