package main

import "fmt"

type User struct {
	Username string
	Password string
}

var users = map[string]string{
	"George": "Gichuru02",
	"User2":  "password2",
}

func authenticate(username, password string) bool {
	storedPassword, ok := users[username]
	if !ok {
		return false
	}
	return storedPassword == password
}

func main() {
	username := "George"
	password := "Gichuru02"

	if authenticate(username, password) {
		fmt.Println("Authentication successful. Granted access")
	} else {
		fmt.Println("Authentication failed. Access denied")
	}
}
