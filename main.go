package main

import (
	"log"

	"github.com/kynmh69/go-rds/domain"
	"github.com/kynmh69/go-rds/infrastructure/database"
	"github.com/kynmh69/go-rds/usecases"
)

func main() {
	con, err := database.NewSQLFactory(database.InstanceMySQL)
	if err != nil {
		log.Fatalln(err)
	}
	con.GetConnection().AutoMigrate(&domain.User{})
	usecases.CreateUser(con)
	user := usecases.FindByUsername(con, "test_user_1")
	log.Println(user)
}
