package initialize

import (
	"fmt"
	"github.com/yushengguo557/twitter-space/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() {
	dbCfg := global.App.Config.Database // 数据库配置
	//dsn := "nft_dev:k4r35kbMa7Jr3dmc@tcp(180.184.138.237:3306)/nft_dev?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "nft:FcPBSEzftrRKDLFF12#@tcp(180.184.170.180:3306)/nft?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
		dbCfg.Charset)
	var err error

	// 连接数据库
	global.App.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("connect database, err: %w", err))
	}
	log.Printf("数据库: %s(%s:%d) 已连接\n",
		dbCfg.DBName,
		dbCfg.Host,
		dbCfg.Port)
}
