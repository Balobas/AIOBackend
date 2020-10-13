package api

import (
	"../database"
	"../modules/socialNetworks"
	"encoding/json"
	"../models"
)

func DisconnectSocialNetworkAccount(accountJSON []byte) string {
	var account models.SocialNetworkAccount
	err := json.Unmarshal(accountJSON, &account)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	err = socialNetworks.DisconnectSocialNetworkAccount(&database.DATABASE, account)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{\"Status\" : \"OK\"}"
}