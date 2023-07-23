package main

import (
	"time"
	"twitter-space/db"
	"twitter-space/global"
	"twitter-space/tw"
)

const Period time.Duration = time.Hour * 1

func init() {
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
