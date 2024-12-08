package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func CreateDb() {
	DB, _ = sql.Open("sqlite3", "./db/cs.db")

	// Create a table
	createTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
		password TEXT
    );
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nom TEXT,
		adresse TEXT,
		siret TEXT
	);
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER,
		produit_id INTEGER,
		quantite INTEGER,
		date TEXT,
		status TEXT
	);
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		quantity INTEGER,
		price INTEGER,
		name TEXT,
		date TEXT,
		category TEXT,
		emplacement TEXT
	);
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);`
	if _, err := DB.Exec(createTable); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
