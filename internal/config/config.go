package config

import (
	"os"

	"github.com/gabehamasaki/encurtago/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ENV         string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB          *database.Queries
}

func NewConfig() *Config {
	return &Config{
		ENV:         os.Getenv("ENV"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
	}
}

func (c *Config) SetDB(db *database.Queries) {
	c.DB = db
}
