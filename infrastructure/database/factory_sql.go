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

func NewSQLFactory(instance int) (database.Connection, error) {
	switch instance {
	case InstanceMySQL:
		return NewMySQLConnection(), nil
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
