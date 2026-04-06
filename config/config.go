package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type TConfig struct {
	PORT         string
	DATABASE_URL string
	JWT_SECRET   string
	HASH_SALT    int
}

func LoadConfig(envFile string) *TConfig {

	err := godotenv.Load(envFile)

	if err != nil {
		log.Println("No se encontro el archivo .env, se usaran las variables de entorno del sistema")
	}

	return &TConfig{
		PORT:         getEnv("PORT", "8080"),
		DATABASE_URL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb"),
		JWT_SECRET:   getEnv("JWT_SECRET", "mysecret"),
		HASH_SALT:    func() int { v, _ := strconv.Atoi(getEnv("HASH_SALT", "10")); return v }(),
	}

}

func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	return val
}
