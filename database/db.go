// database/db.go
package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    host := os.Getenv("DB_HOST")   // "db" จาก docker-compose
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to open DB connection:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    log.Println("Connected to PostgreSQL!")
}

func Migrate() {
    query := `
    CREATE TABLE IF NOT EXISTS contacts (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
		company VARCHAR(100),
        email VARCHAR(100),
        phone VARCHAR(50),
        address TEXT
    );

    CREATE TABLE IF NOT EXISTS recommendations (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100)
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatal("Failed to run migrations:", err)
    }

    log.Println("Migrations complete")
}




// package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

// func Connect() {
// 	connStr := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s sslmode=disable",
// 		os.Getenv("DB_HOST"),
// 		os.Getenv("DB_USER"),
// 		os.Getenv("DB_PASSWORD"),
// 		os.Getenv("DB_NAME"),
// 	)

// 	var err error
// 	DB, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal("Cannot connect to DB:", err)
// 	}
// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal("Cannot ping DB:", err)
// 	}
// 	log.Println("Connected to PostgreSQL")
// }
