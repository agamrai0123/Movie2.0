package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Printf("DB connection established")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("Error closing DB connection: %v", err)
	}
	log.Printf("DB connection closed")
	db = nil
}

func GetActiveConnections(db *sql.DB) (int, error) {
	var activeConnections int
	query := "SELECT COUNT(*) FROM pg_stat_activity WHERE state = 'active'"
	err := db.QueryRow(query).Scan(&activeConnections)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return -1, err
	}
	return activeConnections, nil
}
