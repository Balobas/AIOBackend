package users

import (
	"../../../data"
	"../../database"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(database database.Database, user data.User, password string) (data.UID, error) {
	uid, err := PutUser(database, user)
	if err != nil {
		return "", err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	userPassword := data.UserPassword {
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