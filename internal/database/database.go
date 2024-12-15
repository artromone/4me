package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

type Database struct {
    Conn *sql.DB
}

func NewDatabase() *Database {
    // Database connection parameters
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    // Connection string
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Open connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    // Verify connection
    if err = db.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    return &Database{Conn: db}
}

func (d *Database) Close() {
    d.Conn.Close()
}

func (d *Database) Migrate() error {
    // SQL to create initial tables
    _, err := d.Conn.Exec(`
        CREATE TABLE IF NOT EXISTS groups (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            description TEXT
        );

        CREATE TABLE IF NOT EXISTS lists (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            group_id INTEGER REFERENCES groups(id),
            description TEXT
        );

        CREATE TABLE IF NOT EXISTS tasks (
            id SERIAL PRIMARY KEY,
            title VARCHAR(200) NOT NULL,
            description TEXT,
            status VARCHAR(20) DEFAULT 'pending',
            list_id INTEGER REFERENCES lists(id),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            due_date TIMESTAMP
        );

        CREATE TABLE IF NOT EXISTS collaborators (
            id SERIAL PRIMARY KEY,
            list_id INTEGER REFERENCES lists(id),
            user_id INTEGER,
            role VARCHAR(50) DEFAULT 'viewer'
        );
    `)

    return err
}
