package data

import "github.com/dgrijalva/jwt-go"

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

type UserPassword struct {
	UserUID UID `json:"userUid"`
	Password string `json:"password"`
}

type SocialNetwork struct {
	UID       UID    `json:"uid"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

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

type UserSettings struct {
	UserUID UID `json:"userUid"`
}

type AccessToken struct {
	UserUid UID `json:"userUid"`
	jwt.StandardClaims
}