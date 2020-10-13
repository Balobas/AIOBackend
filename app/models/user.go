package models

import "../../data"

//Аккаунт пользователя в приложении
type User struct {
	UID      data.UID    `json:"uid"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Login    string `json:"login"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Country  string `json:"country"`
	Location string `json:"location"`
	Link     string `json:"link"`
	IsArchived bool `json:"isArchived"`
	Token string `json:"token"`
}