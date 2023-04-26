package config

import "time"

type RoutinePoolConfig struct {
	PoolSize       int           `yaml:"poolSize"`
	ExpiryDuration time.Duration `yaml:"expiryDuration"`
}
