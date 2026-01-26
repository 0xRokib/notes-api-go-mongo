package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load env...")
	}
	mongoURI, err := ExtractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}
	mongoDB, err := ExtractEnv("MONGO_DB")
	if err != nil {
		return Config{}, err
	}
	port, err := ExtractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil

}

func ExtractEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("Missing req env...")
	}
	return val, nil
}
