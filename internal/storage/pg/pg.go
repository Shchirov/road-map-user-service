package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"road-map-user-server/internal/config"
)

func Dial(config *config.DataBaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.PostgresHost,
		config.PgUser,
		config.PgPassword,
		config.PostgresDb,
		config.PostgresPort)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("todo")
	}
	return db
}
