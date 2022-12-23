package main

import (
	"gotodo/database"
	"log"
	"os"
)

func main() {
	username := os.Getenv("username")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbName := os.Getenv("db")

	credsDatabase := map[string]interface{}{
		"username": username,
		"password": password,
		"host":     host,
		"port":     port,
		"dbName":   dbName,
	}

	_, err := database.Connection(credsDatabase)
	if err != nil {
		log.Fatal(err.Error())
	}

}
