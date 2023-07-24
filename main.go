package main

import (
	"github.com/yushengguo557/twitter-space/initialize"
)

func init() {
	initialize.Init() // 初始化
}

func main() {
	// 开始定时任务
	Corn()
	// 阻塞主线程
	var wait = make(chan struct{})
	wait <- struct{}{}
}
