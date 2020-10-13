package socialNetworks

import (
	"../../../data"
	"../../database"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"../../models"
)

//Добавление аккаунта соц сети в список аккаунтов пользователя
func ConnectSocialNetworkAccount(database database.Database, account models.SocialNetworkAccount, userUID data.UID) (data.UID, error) {
	if !userUID.IsCorrect() {
		return "", errors.New("User uid is not correct ")
	}
	if account.SocialNetworkUid == "" || !account.SocialNetworkUid.IsCorrect() {
		return "", errors.New("Invalid social network uid format")
	}
	_, isFound, err := GetSocialNetwork(database, account.SocialNetworkUid)
	if err != nil {
		return "", err
	}
	if !isFound {
		return "", errors.New("Social network is not supported ")
	}
	account.UserUID = userUID
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	account.UID = data.UID(uid.String())
	err = database.Set(string(account.UID), account)
	if err != nil {
		return "", err
	}
	return account.UID, nil
}