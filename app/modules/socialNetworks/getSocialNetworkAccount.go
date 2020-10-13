package socialNetworks

import (
	"../../../data"
	"../../database"
	"../../models"
	"encoding/json"
	"github.com/pkg/errors"
)

//Получение аккаунта в социальной сети по id
func GetSocialNetworkAccount(database database.Database, uid data.UID) (models.SocialNetworkAccount, bool, error) {
	if !uid.IsCorrect() {
		return models.SocialNetworkAccount{}, false, errors.New("Invalid uid format ")
	}
	fieldsMap, err := database.Get(string(uid))
	if err != nil {
		return models.SocialNetworkAccount{}, false, err
	}
	if fieldsMap == nil {
		return models.SocialNetworkAccount{}, false, nil
	}
	var account models.SocialNetworkAccount
	bytes, err  := json.Marshal(fieldsMap)
	if err != nil {
		return models.SocialNetworkAccount{}, false, err
	}
	err = json.Unmarshal(bytes, &account)
	if err != nil {
		return models.SocialNetworkAccount{}, false, err
	}
	return account, true, nil
}
