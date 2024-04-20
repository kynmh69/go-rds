package database

type Connection interface {
	ConnectDb()
	ConnectWriteDb()
	ConnectReadDb()
}
