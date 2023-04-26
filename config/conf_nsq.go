package config

type NsqConfig struct {
	NsqdAddr       string `yaml:"nsqdAddr"`
	NsqlookupdAddr string `yaml:"nsqlookupdAddr"`
}
