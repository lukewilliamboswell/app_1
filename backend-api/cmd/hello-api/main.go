package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"hello-world-api/internal/api"
	"log"
	"os"
)

func main() {
	// load environment variables for secured data
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	port := os.Getenv("PORT")
	
	// validate data
	if dbHost == "" || dbPort == "" || dbUsername == "" || dbPassword == "" {
		log.Fatal("DB host, port, username and password required")
	}
	
	// default port
	if port == "" {
		port = "5300"
	}
	
	// open database connection
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			dbHost, dbPort, "hello_db", dbUsername, dbPassword))
	if err != nil {
		log.Fatal(err)
	}
	// close db connection upon application closes
	defer db.Close()
	
	// ping the db to make sure it is alive and working
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	// initiate our API
	app, err := api.NewApi(db)
	if err != nil {
		log.Fatal(err)
	}
	
	// run api
	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
