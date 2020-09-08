package AIO_API

/*
В данном файле собраны функции API, которые доступны из вне.
При создании ответа на любой запрос пользователя к серверу, действия, указанные в запросе не вызываются на прямую,
прежде вызывается соответствующая функция API.
 */


import (
	"../app/database"
	"../app/modules/socialNetworks"
	"../app/modules/users"
	"../data"
	"encoding/json"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

//all functions returns JSON
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

func putUser(userJSON []byte) string {
	var user data.User
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

func CreateUser(userJSON []byte) string {
	return putUser(userJSON)
}

func UpdateUser(userJSON []byte) string {
	return putUser(userJSON)
}

func DeleteUser(userUid data.UID) string {
	err := users.DeleteUser(&database.DATABASE, userUid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{\"Status\" : \"OK\"}"
}

func BackUpUser(userUid data.UID) string {
	err := users.BackupUser(&database.DATABASE, userUid)
	if err != nil {
		return "{\"Error\" : \"" + err.Error() + "\"}"
	}
	return "{\"Status\" : \"OK\"}"
}

func PutSocialNetwork(networkMap map[string][]string) string {
	var network data.SocialNetwork
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

func ConnectSocialNetworkAccount(accountMap map[string][]string, userUid data.UID) string {
	var account data.SocialNetworkAccount
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

func DisconnectSocialNetworkAccount(accountJSON []byte) string {
	var account data.SocialNetworkAccount
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
