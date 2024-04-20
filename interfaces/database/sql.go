package database

type Connection interface {
	ConnectDb()
	connectWriteDb()
	connectReadDb()
}
