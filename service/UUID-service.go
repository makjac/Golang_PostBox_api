package service

import (
	"strings"

	"github.com/google/uuid"
)

func generateUUID(hypen bool) string {
	uuid := uuid.New()
	if hypen {
		return uuid.String()
	}
	return strings.Replace(uuid.String(), "-", "", -1)
}
