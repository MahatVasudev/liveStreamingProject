package config

import "os"

type DBConfig struct {
	DBName     string
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
}

var DBEnvs = DBConfigInstance()

func DBConfigInstance() *DBConfig {
	return &DBConfig{
		DBName:     getEnv("DBName", "livego"),
		Port:       getEnv("DBPort", "5432"),
		DBUser:     getEnv("DBUser", "postgres"),
		DBPassword: getEnv("DBPassword", "root"),
		DBHost:     getEnv("DBHost", "127.0.0.1"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
