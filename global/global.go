package global

import (
	"twitter-space/config"
	"twitter-space/tw"

	"gorm.io/gorm"
)

type Application struct {
	Config        config.Config
	DB            *gorm.DB
	TwitterClient *tw.TwitterClient
}

var App = new(Application)

var DB *gorm.DB

var Client *tw.TwitterClient
