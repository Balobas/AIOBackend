package api

import (
	"../database"
	"../modules/socialNetworks"
	"../models"
)

func PutSocialNetwork(networkMap map[string][]string) string {
	var network models.SocialNetwork
	err := decoder.Decode(&network, networkMap)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	uid, err := socialNetworks.PutSocialNetwork(&database.DATABASE, network)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{ \"uid\" : \"" + string(uid) + "\"}"
}
