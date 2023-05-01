package config

type SystemConfig struct {
	Env       string    `yaml:"env"`
	LogConfig LogConfig `yaml:"logConfig"`
}
