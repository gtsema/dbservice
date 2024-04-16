package db

type User struct {
	Id     int
	ChatId string
	Name   string
	Hash   string
	Salt   string
}
