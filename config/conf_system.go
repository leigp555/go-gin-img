package config

type SystemConfig struct {
	Host      string    `yaml:"host"`
	Port      string    `yaml:"port"`
	Env       string    `yaml:"env"`
	LogConfig LogConfig `yaml:"logConfig"`
}
