package api

func UpdateUser(userJSON []byte) string {
	return putUser(userJSON)
}
