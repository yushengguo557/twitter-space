package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Twitter   Twitter   `yaml:"Twitter"`
	Database  Database  `yaml:"Database"`
	TimedTask TimedTask `yaml:"TimedTask"`
	Limiter   Limiter   `yaml:"Limiter"`
	Proxy     Proxy     `yaml:"Proxy"`
}

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
