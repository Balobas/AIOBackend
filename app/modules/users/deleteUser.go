package users

import (
	"../../../data"
	"../../database"
	"encoding/json"
	"github.com/pkg/errors"
)

//Удаление пользователя по uid. Аккаунт пользователя не удаляется, лишь помечается как удаленный
func DeleteUser(database database.Database, uid data.UID) error {
	if string(uid) == "" {
		return errors.New("Empty user uid")
	}
	if !uid.IsCorrect() {
		return errors.New("Invalid uid format ")
	}
	result, err := database.Get(string(uid))
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("user not found")
	}
	user := data.User{}
	bytes, _ := json.Marshal(result)
	if json.Unmarshal(bytes, &user) != nil {
		return errors.New("Unmarshalling error")
	}
	user.IsArchived = true
	return database.Set(string(user.UID), user)
}
