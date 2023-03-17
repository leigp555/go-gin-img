package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Token  Token  `yaml:"token"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Gin    Gin    `yaml:"gin"`
}
