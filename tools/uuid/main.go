package uuid

import (
	"github.com/google/uuid"
)

const Default = "00000000-0000-0000-0000-000000000000"

func Generate() string {
	return uuid.New().String()
}
