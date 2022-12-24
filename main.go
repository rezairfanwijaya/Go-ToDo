package main

import (
	"fmt"
	"gotodo/database"
	"gotodo/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")

	credsDatabase := map[string]interface{}{
		"username": username,
		"password": password,
		"host":     host,
		"port":     port,
		"dbName":   dbName,
	}

	db, err := database.Connection(credsDatabase)
	if err != nil {
		log.Fatal(err.Error())
	}

	// server
	r := gin.Default()

	// router
	router.NewRoute(db, r)

	// start server
	address := fmt.Sprintf(":%v", 3030)
	if err := r.Run(address); err != nil {
		log.Fatal(err)
	}
}
