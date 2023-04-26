package config

import "time"

type SocketConfig struct {
	Addr              string        `yaml:"addr"`
	ReadBufferSize    int           `yaml:"readBufferSize"`
	WriteBufferSize   int           `yaml:"writeBufferSize"`
	HandshakeTimeout  time.Duration `yaml:"handshakeTimeout"`
	EnableCompression bool          `yaml:"enableCompression"`
}
