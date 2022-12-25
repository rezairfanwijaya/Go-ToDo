package database

import (
	"errors"
	"fmt"

	"gotodo/activity"
	"gotodo/todo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(creds map[string]interface{}) (*gorm.DB, error) {
	username := creds["username"]
	password := creds["password"]
	host := creds["host"]
	port := creds["port"]
	dbName := creds["dbName"]

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		msgErr := fmt.Sprintf("ERR CONNECTION : %v", err.Error())
		return db, errors.New(msgErr)
	}

	db.AutoMigrate(activity.Activity{}, todo.Todo{})

	return db, nil
}
