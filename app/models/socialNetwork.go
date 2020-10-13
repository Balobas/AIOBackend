package models

import "../../data"

//Социальная сеть
type SocialNetwork struct {
	UID       data.UID    `json:"uid"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}