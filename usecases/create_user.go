package usecases

import (
	"github.com/kynmh69/go-rds/domain"
	"github.com/kynmh69/go-rds/infrastructure/encrypt"
	"github.com/kynmh69/go-rds/interfaces/database"
)

func CreateUser(con database.Connection) {
	password, _ := encrypt.HashPassword("password")
	con.GetConnection().Create(&domain.User{Username: "test_user_1", Password: password})
}
