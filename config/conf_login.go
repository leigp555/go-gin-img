package config

type github struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	RedirectUri  string `yaml:"redirectUri"`
	Scope        string `yaml:"scope"`
}
type google struct {
	RedirectUri string `yaml:"redirectUri"`
	Scope       string `yaml:"scope"`
}

type LoginConfig struct {
	Github github `yaml:"github"`
	Google google `yaml:"google"`
}
