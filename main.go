package main

import (
	"log"

	"github.com/kynmh69/go-rds/domain"
	"github.com/kynmh69/go-rds/infrastructure/database"
)

func main() {
	con, err := database.NewSQLFactory(database.InstanceMySQL)
	if err != nil {
		log.Fatalln(err)
	}
	con.GetConnection().AutoMigrate(&domain.User{})
}
