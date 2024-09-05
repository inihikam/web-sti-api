package config

import (
	"os"
)

type Config struct {
	TaBaseURL string
	AlumniBaseURL string
	BkBaseURL string
	KpBaseURL string
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort string
}

func LoadConfig() *Config {
	return &Config{
		TaBaseURL: getEnv("TA_BASE_URL", "https://bimbingan-online.inihikam.my.id/api"),
		AlumniBaseURL: getEnv("ALUMNI_BASE_URL", "https://alumni-sti.inihikam.my.id/api"),
		BkBaseURL: getEnv("BK_BASE_URL", "https://bimbingan-karir.inihikam.my.id/api"),
		KpBaseURL: getEnv("KP_BASE_URL", "https://kp-sti.inihikam.my.id/api"),
	}
}

func ConnectDB() *Config {
	return &Config{
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName: getEnv("DB_NAME", "web-sti"),
		DBHost: getEnv("DB_HOST", "127.0.0.1"),
		DBPort: getEnv("DB_PORT", "3306"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return fallback
}