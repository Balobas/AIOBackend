package socialNetworks

import (
	"../../models"
	"../../database"
	"github.com/pkg/errors"
)

//Удаление соц сети из списка аккаунтов пользователя
func DisconnectSocialNetworkAccount(database database.Database, account models.SocialNetworkAccount) error {
	uid := account.UID
	if uid == "" {
		return errors.New("Empty uid")
	}
	_, isFound, err := GetSocialNetworkAccount(database, account.UID)
	if err != nil {
		return err
	}
	if !isFound {
		return errors.New("Account not found")
	}
	return database.Delete(string(account.UID))
}