package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
	Email EmailConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type EmailConfig struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error while loading .env file, using default config")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		Email: EmailConfig{
			Email: os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address: os.Getenv("ADDRESS"),
		},
	}

}
