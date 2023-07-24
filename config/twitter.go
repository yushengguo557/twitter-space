package config

type Twitter struct {
	APIHost           string `mapstructure:"api_host" yaml:"api_host" json:"api_host"`
	APIKey            string `mapstructure:"api_key" yaml:"api_key" json:"api_key"`
	APIKeySecret      string `mapstructure:"api_key_secret" yaml:"api_key_secret" json:"api_key_secret"`
	BearerToken       string `mapstructure:"bearer_token" yaml:"bearer_token" json:"bearer_token"`
	AccessToken       string `mapstructure:"access_token" yaml:"access_token" json:"access_token"`
	AccessTokenSecret string `mapstructure:"access_token_secret" yaml:"access_token_secret" json:"access_token_secret"`
}
