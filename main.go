package main

import (
	"twitter-space/config"
	"twitter-space/db"
	"twitter-space/global"
	"twitter-space/tw"
)

func init() {
	config.InitConfig()
	db.InitDB()
	global.Client = tw.NewTwitterClient()
}

func main() {
	// 开始定时任务
	Corn()

	// 阻塞主线程
	var wait = make(chan struct{})
	wait <- struct{}{}
}
