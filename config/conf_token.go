package config

type Token struct {
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime int    `yaml:"expiresTime"`
}
