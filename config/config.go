package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	PostgresConn  string
}

func LoadConfig() *Config {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		log.Fatal("SERVER_ADDRESS is not set")
	}

	postgresConn := os.Getenv("POSTGRES_CONN")
	if postgresConn == "" {
		log.Fatal("POSTGRES_CONN is not set")
	}

	return &Config{
		ServerAddress: serverAddress,
		PostgresConn:  postgresConn,
	}
}
