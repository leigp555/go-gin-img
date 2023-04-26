package config

type LogConfig struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`
	ShowLine   bool   `yaml:"showLine"`
	Level      string `yaml:"level"`
}
