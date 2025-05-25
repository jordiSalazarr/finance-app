package domain

import (
	"github.com/google/uuid"
)

type UUID struct {
	Val string
}

func NewUUID() string {
	return uuid.New().String()
}
