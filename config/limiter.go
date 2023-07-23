package config

type Limiter struct {
	Duration int `yaml:"Duration"`
	Quantity int `yaml:"Quantity"`
}
