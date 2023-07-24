# twitter-space
Use twitter api to get twitter space data and update the database regularly.

# 0x00 Start
```shell
# a.set proxy
# macOS || Linux
export GOPROXY=https://goproxy.cn

# Windows
$env:GOPROXY = "https://goproxy.cn"

# b.download dependencies
go mod tidy

# c.run it
chmod +x ./run.sh
./run.sh
```