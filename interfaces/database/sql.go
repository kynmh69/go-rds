package database

import "gorm.io/gorm"

type Connection interface {
	ConnectDb()
	ConnectWriteDb()
	ConnectReadDb()
	GetWriterCon() *gorm.DB
	GetReadCon() *gorm.DB
}
