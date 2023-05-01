package config

import "time"

type SocketConfig struct {
	Mode              string        `yaml:"mode"`
	Host              string        `yaml:"host"`
	Port              string        `yaml:"port"`
	ReadBufferSize    int           `yaml:"readBufferSize"`
	WriteBufferSize   int           `yaml:"writeBufferSize"`
	HandshakeTimeout  time.Duration `yaml:"handshakeTimeout"`
	EnableCompression bool          `yaml:"enableCompression"`
}
