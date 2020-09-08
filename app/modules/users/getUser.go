package users

import (
	"../../../data"
	"../../database"
	"encoding/json"
	"errors"
)

//Получение пользователя по uid
func GetUser(database database.Database, uid data.UID) (data.User, bool, error) {
	if !uid.IsCorrect() {
		return data.User{}, false, errors.New("Invalid uid format ")
	}
	userMap, err := database.Get(string(uid))
	if err != nil {
		return data.User{}, false, err
	}
	if userMap == nil {
		return data.User{}, false, nil
	}
	var user data.User
	userBytes, err  := json.Marshal(userMap)
	if err != nil {
		return data.User{}, false, err
	}
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return data.User{}, false, err
	}
	return user, true, nil
}