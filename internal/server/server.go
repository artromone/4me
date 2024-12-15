package server

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "github.com/artromone/4me/internal/database"
    "github.com/artromone/4me/pkg/config"
)

// Server represents the application server
type Server struct {
    router *mux.Router
    db     *database.Database
    config *config.Config
}

// NewServer creates and configures a new server instance
func NewServer(cfg *config.Config) *Server {
    // Initialize database
    db := database.NewDatabase()

    // Run migrations
    if err := db.Migrate(); err != nil {
        panic(fmt.Sprintf("Database migration failed: %v", err))
    }

    // Create server instance
    s := &Server{
        router: mux.NewRouter(),
        db:     db,
        config: cfg,
    }

    // Setup middleware and routes
    s.setupMiddleware()
    s.setupRoutes()

    return s
}

// setupMiddleware applies global middleware to the router
func (s *Server) setupMiddleware() {
    // Logging middleware
    s.router.Use(loggingMiddleware)
    
    // Recovery middleware
    s.router.Use(recoveryMiddleware)

    // CORS middleware (optional)
    s.router.Use(corsMiddleware)
}

// setupRoutes configures all API routes
func (s *Server) setupRoutes() {
    // Task routes
    s.router.HandleFunc("/tasks", s.createTask).Methods("POST")
    s.router.HandleFunc("/tasks", s.listTasks).Methods("GET")
    s.router.HandleFunc("/tasks/{id}", s.getTask).Methods("GET")
    s.router.HandleFunc("/tasks/{id}", s.updateTask).Methods("PUT")
    s.router.HandleFunc("/tasks/{id}", s.deleteTask).Methods("DELETE")

    // List routes
    s.router.HandleFunc("/lists", s.createList).Methods("POST")
    s.router.HandleFunc("/lists", s.listLists).Methods("GET")
    
    // Group routes
    s.router.HandleFunc("/groups", s.createGroup).Methods("POST")
    s.router.HandleFunc("/groups", s.listGroups).Methods("GET")
}

// Start begins listening on the specified address
func (s *Server) Start() error {
    // Create server with timeouts
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", s.config.ServerPort),
        Handler:      s.router,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    // Print startup message
    fmt.Printf("Starting server on port %d\n", s.config.ServerPort)

    // Start server
    return srv.ListenAndServe()
}

// Graceful shutdown method (to be implemented)
func (s *Server) Shutdown() error {
    // Close database connection
    s.db.Close()
    return nil
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
