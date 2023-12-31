package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	// dotenv for sensitive information
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host, port, dbname, user, password := os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", host, port, dbname, user, password)

	// connect to database
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connect: %v\n", err))
	}
	defer db.Close()
	log.Println("Connected to database")

	// test my connection
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot ping database")
	}
	log.Println(`Database is ping`)

	// get rows from table
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	query := `insert into users (first_name,last_name) values ($1,$2)`
	_, err = db.Exec(query, "Mike", "Mom")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a row!")

	// get rows from table again
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	stmt := `update users set first_name = $1 where id = $2`
	_, err = db.Exec(stmt, "Terraform", 3)
	if err != nil {
		log.Fatal(err)
	}
	// get rows from table again
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}

	// get one row by id
	query = `select id,first_name,last_name from users where id = $1`
	var (
		firstName, lastName string
		id                  int
	)
	row := db.QueryRow(query, 3)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow return", id, firstName, lastName)

	// delete a row
	query = `delete from users where id = $1`
	_, err = db.Exec(query, 7)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted a row")

	// get rows again
	err = getAllRows(db)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(db *sql.DB) error {
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	var (
		firstName, lastName string
		id                  int
	)
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("The data is", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println(strings.Repeat("-", 50))
	return nil
}
