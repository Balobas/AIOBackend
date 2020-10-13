package api

func CreateUser(userJSON []byte) string {
	return putUser(userJSON)
}
