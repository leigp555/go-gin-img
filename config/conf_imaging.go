package config

type ImagingConfig struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
	Blur   int `yaml:"blur"`
}
