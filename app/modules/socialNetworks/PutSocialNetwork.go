package socialNetworks

import (
	"../../../data"
	"../../database"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func PutSocialNetwork(database database.Database, network data.SocialNetwork) (data.UID, error) {
	if network.UID == "" {
		uid, err := uuid.NewV4()
		if err != nil {
			return "", err
		}
		network.UID = data.UID(uid.String())
	}
	if network.Name == "" || network.ShortName == "" {
		return "", errors.New("Empty fields")
	}
	net, isFound, err := GetSocialNetwork(database, network.UID)
	if err != nil {
		return "", err
	}
	if isFound {
		if net.Name != "" {
			network.Name = net.Name
		}
		if net.ShortName != "" {
			net.ShortName = net.ShortName
		}
	}
	err = database.Set(string(network.UID), network)
	if err != nil {
		return "", err
	}
	return network.UID, nil
}