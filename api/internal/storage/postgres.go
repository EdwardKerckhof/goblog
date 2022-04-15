package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
)

func NewPostgresConnection(config *configs.Config) *gorm.DB {
	connectionInfo := ports.DBConnection{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DBName:   config.DBName,
	}

	db, err := gorm.Open(postgres.Open(dbConnectionInfoToString(connectionInfo)), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %s\n", err.Error())
	} else {
		fmt.Printf("database connection succeeded\n")
	}

	pgDB, _ := db.DB()
	err = pgDB.Ping()
	if err != nil {
		log.Fatalf("unable to ping database: %s\n", err.Error())
	} else {
		fmt.Printf("database ping succeeded\n")
	}

	return db
}

func dbConnectionInfoToString(info ports.DBConnection) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Brussels",
		info.Host,
		info.User,
		info.Password,
		info.DBName,
		info.Port)
}
