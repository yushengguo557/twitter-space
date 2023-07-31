package config

type Tos struct {
	AccessKey  string `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey  string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Endpoint   string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	BucketName string `mapstructure:"bucket_name" json:"bucket_name" yaml:"bucket_name"`
}
