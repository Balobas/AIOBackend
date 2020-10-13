package models

import "../../data"

//Аккаунт пользователя в социальной сети
type SocialNetworkAccount struct {
	UID              data.UID    `json:"uid"`
	SocialNetworkUid data.UID    `json:"socialNetUid"`
	UserUID          data.UID    `json:"userUid"`
	UserKey          string `json:"userKey"`
	Login            string `json:"login"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Link             string `json:"link"`
}
