package configs

import (
	"fmt"
	"go-auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Schema   string
}

var DB *gorm.DB

func InitDB(cfg Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.Schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}
