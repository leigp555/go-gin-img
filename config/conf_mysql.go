package config

type Mysql struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	MaxConn  int    `yaml:"maxConn"`
	MaxOpen  int    `yaml:"maxOpen"`
	DB       string `yaml:"db"`
}
