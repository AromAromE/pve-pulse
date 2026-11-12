package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PVEHost  string
	PVENode  string
	PVEToken string
	Port     string
	CacheTTL time.Duration
	DemoMode bool
}

// Load reads configuration from .env and environment variables.
func Load() (*Config, error) {
	// Attempt to load .env file. If it doesn't exist (e.g. in Docker), we don't fail,
	// because environment variables might be passed directly by the OS/container.
	_ = godotenv.Load()

	demoMode := parseBool(getEnv("DEMO_MODE", "false"))
	cacheTTLSeconds := parseInt(getEnv("CACHE_TTL_SECONDS", "4"))

	cfg := &Config{
		PVEHost:  getEnv("PVE_HOST", ""),
		PVENode:  getEnv("PVE_NODE", "aromarome"),
		PVEToken: getEnv("PVE_TOKEN", ""),
		Port:     getEnv("PORT", "8080"),
		CacheTTL: time.Duration(cacheTTLSeconds) * time.Second,
		DemoMode: demoMode,
	}

	// Fail-fast validation
	if !cfg.DemoMode {
		if cfg.PVEHost == "" {
			return nil, fmt.Errorf("config error: PVE_HOST is required when DEMO_MODE is false")
		}
		if cfg.PVEToken == "" {
			return nil, fmt.Errorf("config error: PVE_TOKEN is required when DEMO_MODE is false")
		}
	}

	return cfg, nil
}

// Helper: Read an env var with a fallback default string
func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultVal
}

// Helper: Convert string to int with a safe fallback
func parseInt(val string) int {
	res, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return res
}

// Helper: Convert string to bool (accepts "1", "t", "T", "true", "TRUE")
func parseBool(val string) bool {
	res, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	return res
}