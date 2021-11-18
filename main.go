package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"rest/database"
)
func main() {
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

}
