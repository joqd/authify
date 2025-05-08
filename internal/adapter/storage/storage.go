package storage

import "gorm.io/gorm"

type PostgresDatabase interface {
	GetDB() *gorm.DB
	AutoMigrate() error
}
