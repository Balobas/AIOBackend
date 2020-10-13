package models

import "../../data"

//Пароль пользователя
type UserPassword struct {
	UserUID data.UID `json:"userUid"`
	Password string `json:"password"`
}