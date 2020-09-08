package data

/*
Структуры
 */

import "github.com/dgrijalva/jwt-go"

//Аккаунт пользователя в приложении
type User struct {
	UID      UID    `json:"uid"`
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

//Пароль пользователя
type UserPassword struct {
	UserUID UID `json:"userUid"`
	Password string `json:"password"`
}

//Социальная сеть
type SocialNetwork struct {
	UID       UID    `json:"uid"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

//Аккаунт пользователя в социальной сети
type SocialNetworkAccount struct {
	UID              UID    `json:"uid"`
	SocialNetworkUid UID    `json:"socialNetUid"`
	UserUID          UID    `json:"userUid"`
	UserKey          string `json:"userKey"`
	Login            string `json:"login"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Link             string `json:"link"`
}

//Настройки аккаунта
type UserSettings struct {
	UserUID UID `json:"userUid"`
}

//Токен доступа
type AccessToken struct {
	UserUid UID `json:"userUid"`
	jwt.StandardClaims
}