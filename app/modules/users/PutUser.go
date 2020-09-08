package users

import (
	"../../../data"
	"../../database"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func PutUser(database database.Database, user data.User) (data.UID, error) {
	var (
		isFound bool
		savedUser data.User
		err error
	)
	if user.UID == "" {
		uid, err := uuid.NewV4()
		if err != nil {
			return "", errors.New("uid generation error")
		}
		user.UID = data.UID(uid.String())
		isFound = false
	} else {
		if !user.UID.IsCorrect() {
			return "", errors.New("Invalid uid")
		}
		savedUser, isFound, err = GetUser(database, user.UID)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
	}
	if !isFound {
		if err := ValidateUser(&user); err != nil {
			return "", err
		}
		userWithSamePhoneNumber, err := database.QueryAllFieldsWithSelector(`phone="` + user.Phone + `"`)
		if err != nil {
			return "", err
		}
		if len(userWithSamePhoneNumber) != 0 {
			return "", errors.New("This phone number already busy ")
		}
		savedUser = user
	} else {
		if user.Name != "" {
			savedUser.Name = user.Name
		}
		if user.FullName != "" {
			savedUser.FullName = user.FullName
		}
		if user.Email != "" {
			savedUser.Email = user.Email
		}
		savedUser.IsArchived = user.IsArchived
		if err := ValidateUser(&savedUser); err != nil {
			return "", err
		}
	}
	return savedUser.UID, database.Set(string(savedUser.UID), savedUser)
}