package models

import (
	"errors"
	"strings"
)

// User fields
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var userList = []User{
	User{ID: 1, Username: "Miguel Villarreal", Password: "venom789"},
}

// IsUserValid checks if the username and password combination is valid
func IsUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

// RegisterNewUser creates a new user with the given username and password
func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := User{ID: uint64(len(userList) + 1), Username: username, Password: password}

	userList = append(userList, u)

	return &u, nil
}

// IsUsernameAvailable checks if the supplied username is available
func IsUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
