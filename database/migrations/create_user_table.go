package main

import (
	"db-connection/database"
	"db-connection/models"
	"log"
)

func CreateUserTable(conn database.DBconnetion) {
	userModel := new(models.User)
	err := conn.AutoMigrate(userModel)
	if err != nil {
		log.Println("impossible migrate User", err)
	}
}
