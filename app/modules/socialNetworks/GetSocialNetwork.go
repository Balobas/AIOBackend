package socialNetworks

import (
	"../../../data"
	"../../database"
	"encoding/json"
	"github.com/pkg/errors"
)

//Получение социальной сети по uid
func GetSocialNetwork(database database.Database, uid data.UID) (data.SocialNetwork, bool,  error) {
	if !uid.IsCorrect() {
		return data.SocialNetwork{}, false, errors.New("Invalid uid format ")
	}
	fieldsMap, err := database.Get(string(uid))
	if err != nil {
		return data.SocialNetwork{}, false, err
	}
	if fieldsMap == nil {
		return data.SocialNetwork{}, false, nil
	}
	var network data.SocialNetwork
	bytes, err  := json.Marshal(fieldsMap)
	if err != nil {
		return data.SocialNetwork{}, false, err
	}
	err = json.Unmarshal(bytes, &network)
	if err != nil {
		return data.SocialNetwork{}, false, err
	}
	return network, true, nil
}