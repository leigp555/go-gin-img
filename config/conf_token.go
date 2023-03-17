package config

type Token struct {
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime string `yaml:"expiresTime"`
}
