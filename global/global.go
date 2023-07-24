package global

import (
	"github.com/yushengguo557/twitter-space/config"
	"github.com/yushengguo557/twitter-space/tw"
	"gorm.io/gorm"
)

type Application struct {
	Config        config.Configuration
	DB            *gorm.DB
	TwitterClient *tw.TwitterClient
}

var App = new(Application)
