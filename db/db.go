package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"twitter-space/global"
)

func InitDB() {
	//dsn := "nft_dev:k4r35kbMa7Jr3dmc@tcp(180.184.138.237:3306)/nft_dev?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "nft:FcPBSEzftrRKDLFF12#@tcp(180.184.170.180:3306)/nft?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("数据库: %s 已连接\n", dsn)
}
