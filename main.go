package main

import (
	"log"
	"net/http"

	"github.com/Henry-Asante/auth_service/controllers"
	"github.com/Henry-Asante/auth_service/database"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)


func init() {
	gotenv.Load()
}

func main() {
	db := database.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
