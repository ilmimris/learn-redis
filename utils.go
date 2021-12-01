package main

import (
	"encoding/json"

	"github.com/ilmimris/learn-redis/domain"
)

// Converts from []byte to a json object according to the User struct.
func toJson(val []byte) domain.User {
	user := domain.User{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}
	return user
}
