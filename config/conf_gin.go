package config

type GinConfig struct {
	Mode      string    `yaml:"mode"`
	LogConfig LogConfig `yaml:"logConfig"`
}
