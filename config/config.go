package config

import (
	"os"
)

type Config struct {
	TaBaseURL string
	AlumniBaseURL string
	BkBaseURL string
	KpBaseURL string
}

func LoadConfig() *Config {
	return &Config{
		TaBaseURL: getEnv("TA_BASE_URL", "https://bimbingan-online.inihikam.my.id/api"),
		AlumniBaseURL: getEnv("ALUMNI_BASE_URL", "https://alumni-sti.inihikam.my.id/api"),
		BkBaseURL: getEnv("BK_BASE_URL", "https://bimbingan-karir.inihikam.my.id/api"),
		KpBaseURL: getEnv("KP_BASE_URL", "https://kp-sti.inihikam.my.id/api"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return fallback
}