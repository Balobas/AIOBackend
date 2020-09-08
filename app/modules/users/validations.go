package users

import (
	"../../../data"
	"errors"
	"regexp"
)

/*
Валидация данных пользователя
 */
func ValidateUser(user *data.User) error {
	if user.Login == "" {
		return errors.New("Error: empty login ")
	}
	if user.Email != "" {
		pattern := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
		if ok, err := regexp.MatchString(pattern, user.Email); !ok || err != nil {
			return errors.New("invalid email")
		}
	}
	if user.Phone == "" {
		return errors.New("Error: empty phone ")
	}
	country, err := ValidatePhoneByLocation(user.Phone)
	if err != nil {
		return err
	}
	user.Country = country
	if user.Name == "" {
		return errors.New("Error: empty name ")
	}
	return nil
}


/*Валидация телефона пользователя. По телефону возвращается локация.
Если регион не поддерживается,то возвращается ошибка.
 */
func ValidatePhoneByLocation(number string) (string, error) {
	for pattern, country := range data.PhonesMap {
		if ok, err := regexp.MatchString(pattern, number); err != nil {
			return "", err
		} else if ok {
			return country, nil
		}
	}
	return "", errors.New("Country is not supported ")
}
