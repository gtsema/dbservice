package dbserver

import (
	"dbservice/config"
	"log"
	"net/http"
)

func Start() {
	http.Handle("/createUser", nil)
	http.Handle("/readUser", nil)
	http.Handle("/updateUser", nil)
	http.Handle("/deleteUser", nil)

	log.Fatal(http.ListenAndServe(config.GetAddr(), nil))
}
