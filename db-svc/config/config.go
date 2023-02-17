package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Conf Config

// Config [Root config structure]
type Config struct {
	Host     string
	Postgres PostgresConfig
}

type PostgresConfig struct {
	ConnectionString string
}

// Override default values with env
func LoadEnv() {
	Conf.Host = getEnv("HOST", Conf.Host)
	Conf.Postgres.ConnectionString = getEnv("POSTGRES_CONNECTION_STRING", Conf.Postgres.ConnectionString)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func LoadConfJson() error {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return err
}
