package socialNetworks

import (
	"../../../data"
	"../../database"
	"github.com/pkg/errors"
)

func DisconnectSocialNetworkAccount(database database.Database, account data.SocialNetworkAccount) error {
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