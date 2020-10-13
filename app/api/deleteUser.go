package api

import (
	"../database"
	"../modules/users"
	"../../data"
)

func DeleteUser(userUid data.UID) string {
	err := users.DeleteUser(&database.DATABASE, userUid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{\"Status\" : \"OK\"}"
}
