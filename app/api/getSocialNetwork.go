package api

import (
	"../../data"
	"../database"
	"../modules/socialNetworks"
	"encoding/json"
)

func GetSocialNetwork(uid data.UID) string {
	network, isFound, err := socialNetworks.GetSocialNetwork(&database.DATABASE, uid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	if !isFound {
		return "{}"
	}
	bytes, err := json.Marshal(network)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return string(bytes)
}
