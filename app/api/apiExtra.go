package api

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"../database"
	"../modules/users"
	"../models"
)

var decoder = schema.NewDecoder()

func putUser(userJSON []byte) string {
	var user models.User
	err := json.Unmarshal(userJSON, &user)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	uid, err := users.PutUser(&database.DATABASE, user)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{ \"uid\" : \"" + string(uid) + "\"}"
}