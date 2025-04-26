package postgres

import (
	"fmt"
	"log"
	"sync"

	"github.com/joqd/authify/internal/adapter/config"
	"github.com/joqd/authify/internal/adapter/storage"
	"github.com/joqd/authify/internal/adapter/storage/postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	DB *gorm.DB
}

var (
	dbInstance *postgresDatabase
	once       sync.Once
)

func NewPostgresDatabase(config *config.Config) storage.PostgresDatabase {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			config.PostgresHost,
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresDB,
			config.PostgresPort,
			config.PostgresSSLMode,
			config.PostgresTimeZone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect database: %v", err)
		}

		dbInstance = &postgresDatabase{DB: db}
	})

	return dbInstance
}

func (pd *postgresDatabase) GetDB() *gorm.DB {
	return pd.DB
}


func (pd *postgresDatabase) AutoMigrate() error {
	return pd.DB.AutoMigrate(
		&model.UserModel{},
		&model.GroupModel{},
		&model.PermissionModel{},
		&model.AccessTokenModel{},
		&model.APIKeyModel{},
		&model.EmailVerificationTokenModel{},
		&model.PasswordResetTokenModel{},
		&model.RefreshTokenModel{},
		&model.UserActivityModel{},
	)
}