package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log"
	"os"
)

func NewPsqlDB() *sql.DB {

	// 1. Define the Datasource name
	// Get the data source name from .env file
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"), os.Getenv("TIMEZONE"))

	fmt.Println(dsn)
	// 2. Open the sql database with driver 'postgres' and datasource name which is just defined
	fmt.Println("Connecting to database...")
	db, err := sql.Open("postgres", dsn)

	// 3. Catch the error (if != nil)
	// Log and exit app if there are any error
	// Only log if anything else
	if err != nil {
		log.Fatalln("Can't connect to database", err)
	} else {
		fmt.Println("Connected to database")
	}

	// 4. Return database which is defined
	return db
}
