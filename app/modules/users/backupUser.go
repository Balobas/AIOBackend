package users

import (
	"../../../data"
	"../../database"
	"github.com/pkg/errors"
)

//Восстановление аккаунта пользователя
func BackupUser(database database.Database, userUid data.UID) error {
	user, isFound, err := GetUser(database, userUid)
	if err != nil {
		return err
	}
	if !isFound {
		return errors.New("User not found ")
	}
	if !user.IsArchived {
		return errors.New("Cant back up not deleted account")
	}
	user.IsArchived = false
	_, err = PutUser(database, user)
	if err != nil {
		return err
	}
	return nil
}