package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnection struct {
	WriteDb *gorm.DB
	ReadDb  *gorm.DB
}

func (con *MySQLConnection) ConnectDb() {
	con.connectReadDb()
	con.connectWriteDb()
}

func (con *MySQLConnection) connectWriteDb() {
	var err error
	config := NewMySQLConf()

	con.WriteDb, err = gorm.Open(mysql.Open(config.ConfigWrite.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open write db")
	}
}

func (con *MySQLConnection) connectReadDb() {
	var err error
	config := NewMySQLConf()

	con.WriteDb, err = gorm.Open(mysql.Open(config.ConfigRead.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalln("cannot open read db")
	}
}
