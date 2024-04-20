package database

import "gorm.io/gorm"

type Connection interface {
	ConnectDb()
	ConnectWriteDb()
	ConnectReadDb()
	GetConnection() *gorm.DB
}
