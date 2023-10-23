package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	c"github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/config"
)

func ConnectDatabase(config c.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	return db, dbErr
}
