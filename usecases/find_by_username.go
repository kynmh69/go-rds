package usecases

import (
	"github.com/kynmh69/go-rds/domain"
	"github.com/kynmh69/go-rds/interfaces/database"
)

func FindByUsername(con database.Connection, username string) *domain.User {
	var user domain.User
	con.GetConnection().First(&user, "username = ?", username)
	return &user
}
