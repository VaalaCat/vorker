package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
