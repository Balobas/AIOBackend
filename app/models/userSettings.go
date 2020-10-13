package models

import "../../data"

//Настройки аккаунта
type UserSettings struct {
	UserUID data.UID `json:"userUid"`

}
