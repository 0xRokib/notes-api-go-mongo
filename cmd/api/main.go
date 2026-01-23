package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	// Connect to the database
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	// Ensure the database connection is closed on exit
	defer func() {
		if err := db.DisConnect(client); err != nil {
			log.Printf("database disconnect error: %v", err)
		}
	}()
	// Start the server
	router := server.NewRouter(database)
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
