package config

import (
	"github.com/luizhirasawa/Go-api.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := "host=localhost user=luiz password=1304 dbname=GOapi port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("psql op error: ", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("psql op error: ", err)
		return nil, err
	}
	return db, nil
}
