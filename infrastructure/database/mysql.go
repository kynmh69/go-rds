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
	WriteDb *gorm.DB
	ReadDb  *gorm.DB
}

func (con *mySQLConnection) ConnectDb() {
	con.ConnectReadDb()
	con.ConnectWriteDb()
}

func (con *mySQLConnection) ConnectWriteDb() {
	var err error
	config := NewMySQLConf()

	con.WriteDb, err = gorm.Open(mysql.Open(config.ConfigWrite.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open write db")
	}
}

func (con *mySQLConnection) ConnectReadDb() {
	var err error
	config := NewMySQLConf()

	con.WriteDb, err = gorm.Open(mysql.Open(config.ConfigRead.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open read db")
	}
}
