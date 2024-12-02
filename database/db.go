package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect establishes a connection to the PostgreSQL database
func Connect() *sql.DB {
	// Supabase connection details
	dbHost := "aws-0-ap-southeast-1.pooler.supabase.com"
	dbPort := "6543"
	dbUser := "postgres.auodbtvdrjrpblxxxxvn"
	dbPassword := "!Bombompokemon3000"
	dbName := "postgres"

	// Construct the PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Open the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Connected to the database successfully.")
	return db
}
