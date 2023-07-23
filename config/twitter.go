package config

type Twitter struct {
	APIKey            string `yaml:"APIKey"`
	APIKeySecret      string `yaml:"APIKeySecret"`
	BearerToken       string `yaml:"BearerToken"`
	AccessToken       string `yaml:"AccessToken"`
	AccessTokenSecret string `yaml:"AccessTokenSecret"`
}
