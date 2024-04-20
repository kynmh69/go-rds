package database

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/kynmh69/go-rds/interfaces/database"
)

const PROTOCOL = "tcp"

type config struct {
	ConfigRead  *mysql.Config
	ConfigWrite *mysql.Config
}

func (c *config) SetConfigWrite() {
	c.ConfigWrite = mysql.NewConfig()

	c.ConfigWrite.Addr = fmt.Sprintf("%s:%s", c.getHostWrite(), c.getPort())
	c.ConfigWrite.User = c.getUserName()
	c.ConfigWrite.Passwd = c.getPasswd()
	c.ConfigWrite.DBName = c.getDbName()
	c.ConfigWrite.Net = c.getProtocol()
}

func (c *config) SetConfigRead() {
	c.ConfigRead = mysql.NewConfig()
	c.ConfigRead.Addr = fmt.Sprintf("%s:%s", c.getHostRead(), c.getPort())
	c.ConfigRead.User = c.getUserName()
	c.ConfigRead.Passwd = c.getPasswd()
	c.ConfigRead.DBName = c.getDbName()
	c.ConfigRead.Net = c.getProtocol()
}

func (c *config) getHostWrite() string {
	hostname := "database_writer"
	if h, ok := os.LookupEnv(database.DB_HOST); ok {
		hostname = h
	}
	return hostname
}
func (c *config) getHostRead() string {
	hostname := "database_reader"
	if h, ok := os.LookupEnv(database.DB_HOST_R); ok {
		hostname = h
	}
	return hostname
}
func (c *config) getUserName() string {
	username := "app"
	if u, ok := os.LookupEnv(database.DB_USER); ok {
		username = u
	}
	return username
}

func (c *config) getPasswd() string {
	password := "passwd"
	if p, ok := os.LookupEnv(database.DB_PASS); ok {
		password = p
	}
	return password
}

func (c *config) getPort() string {
	port := "3306"
	if p, ok := os.LookupEnv(database.DB_PORT); ok {
		port = p
	}
	return port
}

func (c *config) getDbName() string {
	dbName := "3306"
	if d, ok := os.LookupEnv(database.DB_NAME); ok {
		dbName = d
	}
	return dbName
}

func (c *config) getProtocol() string {
	protocol := PROTOCOL
	return protocol
}

func NewMySQLConf() *config {
	c := config{}
	c.SetConfigWrite()
	c.SetConfigRead()
	return &c
}
