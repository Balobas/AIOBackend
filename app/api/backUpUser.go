package api

import (
	"../database"
	"../modules/users"
	"../../data"
)

func BackUpUser(userUid data.UID) string {
	err := users.BackupUser(&database.DATABASE, userUid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{\"Status\" : \"OK\"}"
}
