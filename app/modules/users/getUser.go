package users

import (
	"../../../data"
	"../../database"
	"../../models"
	"encoding/json"
	"errors"
)

//Получение пользователя по uid
func GetUser(database database.Database, uid data.UID) (models.User, bool, error) {
	if !uid.IsCorrect() {
		return models.User{}, false, errors.New("Invalid uid format ")
	}
	userMap, err := database.Get(string(uid))
	if err != nil {
		return models.User{}, false, err
	}
	if userMap == nil {
		return models.User{}, false, nil
	}
	var user models.User
	userBytes, err  := json.Marshal(userMap)
	if err != nil {
		return models.User{}, false, err
	}
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return models.User{}, false, err
	}
	return user, true, nil
}