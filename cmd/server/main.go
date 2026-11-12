package main

import (
	"log"

	"pve-pulse/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Startup failed: %v", err)
	}

	log.Printf("🚀 Config loaded successfully!")
	log.Printf("• Port: %s", cfg.Port)
	log.Printf("• Node: %s", cfg.PVENode)
	log.Printf("• Cache TTL: %v", cfg.CacheTTL)
	log.Printf("• Demo Mode: %t", cfg.DemoMode)
}
