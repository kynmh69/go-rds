package database

import "github.com/go-sql-driver/mysql"

const (
	DB_HOST   = "DB_HOST"
	DB_HOST_R = "DB_HOST_READ"
	DB_PORT   = "DB_PORT"
	DB_USER   = "DB_USER"
	DB_PASS   = "DB_PASS"
	DB_NAME   = "DB_NAME"
	DB_CHAR   = "DB_CHAR"
	DB_LOC    = "DB_LOC"
	DEF_LOC   = "Asia/Tokyo"
)

type Config interface {
	SetConfigWrite() *mysql.Config
	SetConfigRead() *mysql.Config
	getHostWrite() string
	getHostRead() string
	getUserName() string
	getPasswd() string
	getPort() string
	getDbName() string
	getProtocol() string
}
