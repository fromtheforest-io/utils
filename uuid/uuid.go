package uuid

import (
	lib "github.com/google/uuid"
)

func Uuid() string {
	return lib.New().String()
}
