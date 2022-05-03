package service

import (
	"strings"

	"github.com/google/uuid"
)

func generateUUID(hypen bool) string {
	uuid := uuid.New()
	if hypen {
		return uuid.string()
	}
	return strings.Replace(uuid.string(), "-", "", -1)
}
