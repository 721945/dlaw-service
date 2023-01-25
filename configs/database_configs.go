package configs

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

func getConfig() Config {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	name := os.Getenv("DB_NAME")

	return Config{
		User:       user,
		Password:   password,
		Host:       host,
		Port:       port,
		Name:       name,
		DisableTLS: false,
	}
}

func open(cfg Config) (*gorm.DB, error) {
	sslmode := "require"

	if cfg.DisableTLS {
		sslmode = "disable"
	}

	dataSoruce := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslmode)

	return gorm.Open(postgres.Open(dataSoruce), &gorm.Config{})
}

func InitialDatbase() (*gorm.DB, error) {
	return open(getConfig())
}
