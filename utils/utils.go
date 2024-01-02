package utils

import (
	"math/rand"

	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func RandIntFromTo(from, to int) int {
	if from >= 0 {
		return rand.Intn(to+1) + from
	}
	return rand.Intn(to) - from
}
