package config

type Config struct {
	Mysql         MysqlConfig         `yaml:"mysql"`
	Redis         RedisConfig         `yaml:"redis"`
	System        SystemConfig        `yaml:"system"`
	Token         TokenConfig         `yaml:"token"`
	Gin           GinConfig           `yaml:"gin"`
	Login         LoginConfig         `yaml:"login"`
	Email         EmailConfig         `yaml:"email"`
	Imaging       ImagingConfig       `yaml:"imaging"`
	Nsq           NsqConfig           `yaml:"nsq"`
	Socket        SocketConfig        `yaml:"socket"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}
