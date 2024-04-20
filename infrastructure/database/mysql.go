package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func NewMySQLConnection() *mySQLConnection {
	con := mySQLConnection{}
	con.ConnectDb()
	return &con
}

type mySQLConnection struct {
	db     *gorm.DB
	config *config
}

// GetWriterCon implements database.Connection.
func (con *mySQLConnection) GetConnection() *gorm.DB {
	return con.db
}

func (con *mySQLConnection) ConnectDb() {
	con.config = NewMySQLConf()
	con.ConnectWriteDb()
	con.ConnectReadDb()
}

func (con *mySQLConnection) ConnectWriteDb() {
	var err error
	con.db, err = gorm.Open(mysql.Open(con.config.ConfigWrite.FormatDSN()), &gorm.Config{
		Logger: InitLogger(),
	})
	if err != nil {
		log.Fatalln("cannot open write db")
	}
}

func (con *mySQLConnection) ConnectReadDb() {
	con.db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(con.config.ConfigRead.FormatDSN())},
	}))
}
