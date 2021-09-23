package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	config Config
	client *gorm.DB
}

type Option func(db *database)

func New(options ...Option) (*gorm.DB, error) {
	d := &database{}
	for _, v := range options {
		v(d)
	}
	client, err := gorm.Open(mysql.Open(d.config.DSN), &gorm.Config{})
	d.client = client
	return client, err
}

func WithConfig(config Config) Option {
	return func(db *database) {
		db.config = config
	}
}

func WithClient(client *gorm.DB) Option {
	return func(db *database) {
		db.client = client
	}
}
