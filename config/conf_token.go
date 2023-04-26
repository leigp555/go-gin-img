package config

type TokenConfig struct {
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime int    `yaml:"expiresTime"`
}
