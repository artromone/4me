package config

import (
    "os"
    "strconv"
)

type Config struct {
    DatabaseHost     string
    DatabasePort     int
    DatabaseUser     string
    DatabasePassword string
    DatabaseName     string
    ServerPort       int
}

func LoadConfig() *Config {
    port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))

    return &Config{
        DatabaseHost:     getEnv("DB_HOST", "localhost"),
        DatabasePort:     getIntEnv("DB_PORT", 5432),
        DatabaseUser:     getEnv("DB_USER", "taskapp"),
        DatabasePassword: getEnv("DB_PASSWORD", ""),
        DatabaseName:     getEnv("DB_NAME", "taskmanagement"),
        ServerPort:       port,
    }
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getIntEnv(key string, defaultValue int) int {
    valueStr := os.Getenv(key)
    if valueStr == "" {
        return defaultValue
    }
    
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        return defaultValue
    }
    
    return value
}
