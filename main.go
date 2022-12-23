package main

import (
	"gotodo/activity"
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

	db, err := database.Connection(credsDatabase)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := activity.NewActivityRespository(db)
	serv := activity.NewActivityService(repo)

	input := activity.ActivityCreateInput{
		Title: "test",
		Email: "test@gmail.com",
	}

	res, err := serv.CreateActivity(input)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(res)

}
