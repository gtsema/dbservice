package database

type User struct {
	Id     int
	ChatId int
	Name   string
	Hash   string
	Salt   string
}