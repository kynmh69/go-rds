package database

import (
	"errors"

	"github.com/kynmh69/go-rds/interfaces/database"
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)


const (
	InstanceMySQL int = iota
)

func NewDatabaseSQLFactory(instance int) (database.Connection, error) {
	switch instance {
	case InstanceMySQL:
		return 
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}