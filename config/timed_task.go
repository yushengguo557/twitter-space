package config

type TimedTask struct {
	Period int `mapstructure:"period"  yaml:"period" json:"period"`
}
