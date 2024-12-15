package main

import (
    "log"
    "os"

    "github.com/artromone/4me/internal/database"
    "github.com/artromone/4me/internal/server"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize database
    db := database.NewDatabase()
    defer db.Close()

    // Run migrations
    if err := db.Migrate(); err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    // Get server port from environment or use default
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8080"
    }

    // Start server
    srv := server.NewServer(db)
    log.Printf("Starting server on port %s", port)
    if err := srv.Start(":" + port); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
