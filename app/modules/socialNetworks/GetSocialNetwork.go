package socialNetworks

import (
	"../../../data"
	"../../database"
	"encoding/json"
	"github.com/pkg/errors"
	"../../models"
)

//Получение социальной сети по uid
func GetSocialNetwork(database database.Database, uid data.UID) (models.SocialNetwork, bool,  error) {
	if !uid.IsCorrect() {
		return models.SocialNetwork{}, false, errors.New("Invalid uid format ")
	}
	fieldsMap, err := database.Get(string(uid))
	if err != nil {
		return models.SocialNetwork{}, false, err
	}
	if fieldsMap == nil {
		return models.SocialNetwork{}, false, nil
	}
	var network models.SocialNetwork
	bytes, err  := json.Marshal(fieldsMap)
	if err != nil {
		return models.SocialNetwork{}, false, err
	}
	err = json.Unmarshal(bytes, &network)
	if err != nil {
		return models.SocialNetwork{}, false, err
	}
	return network, true, nil
}