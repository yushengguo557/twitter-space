package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yushengguo557/twitter-space/global"
	"log"
)

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.AddConfigPath(".")      // optionally look for config in the working directory
	v.AddConfigPath("/")      // optionally look for config in the working directory
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	err := v.ReadInConfig()   // Find and read the config file
	if err != nil {           // Handle errors reading the config file
		log.Println(err)
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err = v.Unmarshal(&global.App.Config); err != nil {
			log.Println(err)
			panic(fmt.Errorf("unmarshal config: %w", err))
		}
	})
	if err = v.Unmarshal(&global.App.Config); err != nil {
		log.Println(err)
		panic(fmt.Errorf("unmarshal config: %w", err))
	}
}
