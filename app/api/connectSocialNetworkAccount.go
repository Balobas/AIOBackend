package api

import (
	"../database"
	"../modules/socialNetworks"
	"../../data"
	"../models"
)

func ConnectSocialNetworkAccount(accountMap map[string][]string, userUid data.UID) string {
	var account models.SocialNetworkAccount
	err := decoder.Decode(&account, accountMap)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	uid, err := socialNetworks.ConnectSocialNetworkAccount(&database.DATABASE, account, userUid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{ \"uid\" : \"" + string(uid) + "\"}"
}
