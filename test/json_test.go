package test

import (
	"encoding/json"
	"testing"
)

type animal struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Dog struct {
	Animal animal `json:"animal"`
	Gender string `json:"gender"`
}

func TestJson(t *testing.T) {
	dogStr := `{"animal":{"name":"dog","color":"black"},"gender":"male"}`
	dog := new(Dog)
	err := json.Unmarshal([]byte(dogStr), dog)
	if err != nil {
		t.Error(err)
	}
	t.Log(*dog)
}
