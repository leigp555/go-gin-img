package config

import "time"

type Token struct {
	SigningKey  string    `yaml:"signingKey"`
	ExpiresTime time.Time `yaml:"expiresTime"`
}
