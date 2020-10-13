package users

import (
	"../../../data"
	"../../database"
	"golang.org/x/crypto/bcrypt"
	"../../models"
)

/*
Регистрация пользователя. Отличие от добавления пользователя в том,
что при регистрации отдельно создается аккаунт пользователя и отдельно структура пароля.

 */
func RegisterUser(database database.Database, user models.User, password string) (data.UID, error) {
	uid, err := PutUser(database, user)
	if err != nil {
		return "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	userPassword := models.UserPassword {
		UserUID:  uid,
		Password: string(hashedPassword),
	}
	key := "UserPass" + string(uid)
	err = database.Set(key, userPassword)
	if err != nil {
		return uid, err
	}
	return uid, nil
}