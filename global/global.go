package global

import (
	"twitter-space/config"
	"twitter-space/tw"

	"gorm.io/gorm"
)

type App struct {
	Config        config.Config
	DB            *gorm.DB
	TwitterClient *tw.TwitterClient
}

var DB *gorm.DB

var Client *tw.TwitterClient
