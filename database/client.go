package database

import (
	"github.com/jinzhu/gorm"
	"log"
)

var Connector *gorm.DB

func Connect(ConnectionString string) error {
	var err error

	Connector, err = gorm.Open("mysql", ConnectionString)

	if err != nil {
		return err
	}

	log.Println("connection successfully")
	return nil
}
