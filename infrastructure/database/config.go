package database

import (
	"fmt"
	"log"
	"os"
	"time"

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
	c.ConfigWrite.Collation = c.getCharSet()
	c.ConfigWrite.Loc = c.getLocation()
	c.ConfigWrite.ParseTime = true
}

func (c *config) SetConfigRead() {
	c.ConfigRead = mysql.NewConfig()
	c.ConfigRead.Addr = fmt.Sprintf("%s:%s", c.getHostRead(), c.getPort())
	c.ConfigRead.User = c.getUserName()
	c.ConfigRead.Passwd = c.getPasswd()
	c.ConfigRead.DBName = c.getDbName()
	c.ConfigRead.Net = c.getProtocol()
	c.ConfigRead.Collation = c.getCharSet()
	c.ConfigRead.Loc = c.getLocation()
	c.ConfigRead.ParseTime = true
}

func (c *config) getHostWrite() string {
	hostname := "database"
	if h, ok := os.LookupEnv(database.DB_HOST); ok {
		hostname = h
	}
	return hostname
}
func (c *config) getHostRead() string {
	hostname := "database"
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
	dbName := "test"
	if d, ok := os.LookupEnv(database.DB_NAME); ok {
		dbName = d
	}
	return dbName
}

func (c *config) getProtocol() string {
	protocol := PROTOCOL
	return protocol
}

func (c *config) getCharSet() string {
	char := "utf8mb4_unicode_ci"
	if c, ok := os.LookupEnv(database.DB_CHAR); ok {
		char = c
	}
	return char
}

func (c *config) getLocation() *time.Location {
	loc, err := time.LoadLocation(database.DEF_LOC)
	if err != nil {
		log.Fatalln(err)
	}
	if l, ok := os.LookupEnv(database.DB_LOC); ok {
		loc, err = time.LoadLocation(l)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return loc
}

func NewMySQLConf() *config {
	c := config{}
	c.SetConfigWrite()
	c.SetConfigRead()
	return &c
}
