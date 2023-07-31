package config

type Configuration struct {
	Twitter   Twitter   `mapstructure:"twitter" yaml:"twitter" json:"twitter"`
	Database  Database  `mapstructure:"database" yaml:"database" json:"database"`
	TimedTask TimedTask `mapstructure:"timed_task" yaml:"timed_task" json:"timed_task"`
	Limiter   Limiter   `mapstructure:"limiter" yaml:"limiter" json:"limiter"`
	Proxy     Proxy     `mapstructure:"proxy" yaml:"proxy" json:"proxy"`
	Tos       Tos       `mapstructure:"tos" json:"tos" yaml:"tos"`
}
