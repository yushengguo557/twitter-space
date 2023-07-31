package initialize

func Init() {
	InitConfig()        // 初始化配置
	InitDB()            // 初始化数据库
	InitTwitterClient() // 初始化推特客户端
	InitTos()           // 初始化 TOS
}
