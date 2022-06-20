package data

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log"
	"os"
)

type Data struct {
	DB *sql.DB
}

// Init Initial the Database and Connection
func (data *Data) Init() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"), os.Getenv("TIMEZONE"))

	fmt.Println("Connecting to database...")
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalln("Can't connect to database", err)
	} else {
		data.DB = db
		fmt.Println("Connected to database")
	}

	//Migrate here
	fmt.Println("Migrating database...")

	driver, err := postgres.WithInstance(data.DB, &postgres.Config{SchemaName: "public"})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://data/migrations",
		"researchdocument", driver)

	if err != nil {
		log.Fatal("Error creating migration database", err)
	}

	m.Up()

	fmt.Println("Migrate database successfully")
}
