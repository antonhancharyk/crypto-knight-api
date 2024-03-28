package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Connect(connectionString string) error {
	var err error
	db, err := sqlx.Connect("postgres", "user=postgres dbname=yourdatabase sslmode=disable password=yourpassword host=localhost")
	if err != nil {
		log.Fatalln(err)
	}
	return db.Ping()
}

func Get() *sqlx.DB {
	return db
}

func Close() {
	db.Close()
}

// Test the connection to the database
// if err := db.Ping(); err != nil {
// 	log.Fatal(err)
// } else {
// 	log.Println("Successfully Connected")
// }
