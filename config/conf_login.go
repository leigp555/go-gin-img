package config

type Login struct {
	Github struct {
		ClientId     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
	} `yaml:"github"`
}
