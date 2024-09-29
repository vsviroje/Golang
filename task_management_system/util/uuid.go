package util

import (
	uuid "github.com/satori/go.uuid"
)

func GenerateUUIDPk() string {
	id := uuid.NewV4()
	return id.String()
}
