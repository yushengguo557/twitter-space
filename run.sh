export GOPROXY=https://goproxy.cn # 设置代理
go mod tidy # 下载依赖
go build -o twitter-space # 打包成可执行文件
./twitter-space # 运行
