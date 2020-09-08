package data

import (
	uuid "github.com/satori/go.uuid"
)

type UID string

func (uid UID) IsCorrect() bool {
	_, err := uuid.FromString(string(uid))
	if err != nil {
		return false
	}
	return true
}
