package initialize

import (
	"fmt"
	"github.com/yushengguo557/twitter-space/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() {
	//dsn := "nft_dev:k4r35kbMa7Jr3dmc@tcp(180.184.138.237:3306)/nft_dev?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "nft:FcPBSEzftrRKDLFF12#@tcp(180.184.170.180:3306)/nft?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		global.App.Config.Database.Username,
		global.App.Config.Database.Password,
		global.App.Config.Database.Host,
		global.App.Config.Database.Port,
		global.App.Config.Database.DBName,
		global.App.Config.Database.Charset)
	var err error
	global.App.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect database, err: %w", err))
	}
	log.Printf("数据库: %s(%s:%d) 已连接\n",
		global.App.Config.Database.DBName,
		global.App.Config.Database.Host,
		global.App.Config.Database.Port)
}
