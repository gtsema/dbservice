package config

var url = "0.0.0.0"
var port = "8080"

func GetAddr() string {
	return url + ":" + port
}
