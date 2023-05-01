package config

type GinConfig struct {
	Mode      string    `yaml:"mode"`
	Host      string    `yaml:"host"`
	Port      string    `yaml:"port"`
	LogConfig LogConfig `yaml:"logConfig"`
}
