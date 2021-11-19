package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"rest/controller"
	"rest/database"
	"rest/entity"
)

func main() {

	ConnectToDb()

	log.Println("http server run at port 5000")

	router := mux.NewRouter()
	HandleRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

// routing
func HandleRoutes(Route *mux.Router)  {
	// store post
	Route.HandleFunc("/store",controller.Store).Methods("POST")
	// get special post
	Route.HandleFunc("/post/{id}",controller.Show).Methods("GET")
}

// connect to database
func ConnectToDb()  {
	config := database.Conf{
		ServerName: "localhost:3306",
		User:       "daniel",
		Pass:       "Password123#@!",
		Database:   "weblog",
	}

	ConnectStr := database.Connection(config)

	err := database.Connect(ConnectStr)

	if err != nil {
		panic(err.Error())
	}

	// finally migrate table in database
	database.Migration(&entity.Post{})
}

