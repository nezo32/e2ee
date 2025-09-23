package config

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func initDbConfig() *dbConfig {
	return &dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_NAME"),
	}
}

func (cfg *dbConfig) SetupDatabase(dbType DBType) *gorm.DB {
	log.Info("Connecting to database")
	db, err := gorm.Open(dbType.dialector(cfg), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database", "error", err)
	}

	log.Info("Database connection established")

	return db
}

type DBType interface {
	dialector(cfg *dbConfig) gorm.Dialector
	dsn(cfg *dbConfig) string
}

type PostgresDB struct{}

func (db PostgresDB) dsn(cfg *dbConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
}
func (db PostgresDB) dialector(cfg *dbConfig) gorm.Dialector {
	dsn := db.dsn(cfg)
	return postgres.New(postgres.Config{
		DSN: dsn,
	})
}
