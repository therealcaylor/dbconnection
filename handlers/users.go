package handlers

import (
	"db-connection/database"
	"db-connection/models"
	"log"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (t *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbConnection, err := database.GetDbConnection()
	if err != nil {
		log.Println("error", err)
	}
	if r.Method == http.MethodGet {
		GetAllUsers(dbConnection)
	}
}

type arrUser = []*models.User

func GetAllUsers(conn database.DBconnetion) {
	var user []models.User
	conn.Find(&user)
	log.Println(user)

}
