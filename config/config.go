package config

import (
	"fmt"
	"twitter-space/global"

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
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(".")      // optionally look for config in the working directory
	err := v.ReadInConfig()   // Find and read the config file
	if err != nil {           // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err = v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
}
