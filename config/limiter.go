package config

type Limiter struct {
	Duration int `mapstructure:"duration" yaml:"duration" json:"duration"`
	Quantity int `mapstructure:"quantity" yaml:"quantity" json:"quantity"`
}
