package socialNetworks

import (
	"../../../data"
	"../../database"
	"encoding/json"
	"github.com/pkg/errors"
)

//Получение аккаунта в социальной сети по id
func GetSocialNetworkAccount(database database.Database, uid data.UID) (data.SocialNetworkAccount, bool, error) {
	if !uid.IsCorrect() {
		return data.SocialNetworkAccount{}, false, errors.New("Invalid uid format ")
	}
	fieldsMap, err := database.Get(string(uid))
	if err != nil {
		return data.SocialNetworkAccount{}, false, err
	}
	if fieldsMap == nil {
		return data.SocialNetworkAccount{}, false, nil
	}
	var account data.SocialNetworkAccount
	bytes, err  := json.Marshal(fieldsMap)
	if err != nil {
		return data.SocialNetworkAccount{}, false, err
	}
	err = json.Unmarshal(bytes, &account)
	if err != nil {
		return data.SocialNetworkAccount{}, false, err
	}
	return account, true, nil
}
