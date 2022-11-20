package main

import (
	"db-connection/database"
	"log"
)

func main() {
	conn, err := database.GetDbConnection()
	if err != nil {
		log.Println(err)
	}
	CreateUserTable(conn)
}
