package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection() *mySQLConnection {
	con := mySQLConnection{}
	con.ConnectDb()
	return &con
}

type mySQLConnection struct {
	writeDb *gorm.DB
	readDb  *gorm.DB
}

// GetReadCon implements database.Connection.
func (con *mySQLConnection) GetReadCon() *gorm.DB {
	return con.readDb
}

// GetWriterCon implements database.Connection.
func (con *mySQLConnection) GetWriterCon() *gorm.DB {
	return con.writeDb
}

func (con *mySQLConnection) ConnectDb() {
	con.ConnectReadDb()
	con.ConnectWriteDb()
}

func (con *mySQLConnection) ConnectWriteDb() {
	var err error
	config := NewMySQLConf()

	con.writeDb, err = gorm.Open(mysql.Open(config.ConfigWrite.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open write db")
	}
}

func (con *mySQLConnection) ConnectReadDb() {
	var err error
	config := NewMySQLConf()

	con.readDb, err = gorm.Open(mysql.Open(config.ConfigRead.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open read db")
	}
}
