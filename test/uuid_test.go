package test

import (
	"github.com/google/uuid"
	"testing"
)

func generateUuid() (id string) {
	id = uuid.New().String()
	return
}

func TestUUid(t *testing.T) {
	t.Log(generateUuid())
}
