package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"rest/entity"
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

func Migration(table *entity.Post)  {
	Connector.AutoMigrate(&table)

	log.Println("add successfully migrate")
}
