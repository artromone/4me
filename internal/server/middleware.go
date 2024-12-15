package server

import (
    "log"
    "net/http"
    "time"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Log the request
        log.Printf(
            "%s %s %s", 
            r.Method, 
            r.RequestURI, 
            time.Since(start),
        )

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}

func recoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic: %v", err)
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()

        next.ServeHTTP(w, r)
    })
}

func (s *Server) applyMiddleware() {
    s.router.Use(loggingMiddleware)
    s.router.Use(recoveryMiddleware)
}
