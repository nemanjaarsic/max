package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Conf Config

// Config [Root config structure]
type Config struct {
	API_port           string
	Host               string
	DatabaseService    DatabaseSvcConfig
	IdempotencyService IdempotencySvcConfig
}

type DatabaseSvcConfig struct {
	Host string
}

type IdempotencySvcConfig struct {
	Host string
}

// Override default values with env
func LoadEnv() {
	Conf.API_port = getEnv("API_PORT", Conf.API_port)
	Conf.Host = getEnv("HOST", Conf.Host)
	Conf.DatabaseService.Host = getEnv("DATABASE_SERVICE_HOST", Conf.DatabaseService.Host)
	Conf.IdempotencyService.Host = getEnv("IDEMPOTENCY_SERVICE_HOST", Conf.IdempotencyService.Host)
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
