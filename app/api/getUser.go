package api

import (
	"../database"
	"../modules/users"
	"../../data"
	"encoding/json"
)

func GetUser(uidMap map[string][]string) string {
	uid := data.UID(uidMap["uid"][0])
	user, isFound, err := users.GetUser(&database.DATABASE, uid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	if !isFound {
		return "{}"
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return string(bytes)
}
