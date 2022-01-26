package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var database *sql.DB

// InitializeDatabase ...
// Initializes database connection
// for API
func InitializeDatabase() {
	connectionString := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	database = db
}

// GetConnection ...
// Returns a transaction valid for
// 1 db connection.
// Once db interaction is complete
// call tx.Commit() to return the
// connection to the pool
func GetConnection() *sql.Tx {
	tx, err := database.Begin()

	if err != nil {
		panic(err)
	}

	return tx
}

// CloseConnection ...
// Closes connection to db
func CloseConnection(tx *sql.Tx) {
	err := tx.Commit()

	if err != nil {
		panic(err)
	}
}

// PrepareStatement ...
// Prepares statement and
// returns executable stmt object
func PrepareStatement(tx *sql.Tx, query string) *sql.Stmt {
	stmt, err := tx.Prepare(query)
	if err != nil {
		panic(err)
	}

	return stmt
}
