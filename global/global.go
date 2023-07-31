package global

import (
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"github.com/yushengguo557/twitter-space/config"
	"github.com/yushengguo557/twitter-space/tw"
	"gorm.io/gorm"
)

type Application struct {
	Config        config.Configuration
	DB            *gorm.DB
	TwitterClient *tw.TwitterClient
	Tos           *tos.ClientV2
}

var App = new(Application)
