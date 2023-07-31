package initialize

import (
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"github.com/yushengguo557/twitter-space/global"
	"log"
)

// InitTos 初始化 TOS
func InitTos() {
	var err error
	endpoint := global.App.Config.Tos.Endpoint
	region := global.App.Config.Tos.Region
	accessKey := global.App.Config.Tos.AccessKey
	secretKey := global.App.Config.Tos.SecretKey
	global.App.Tos, err = tos.NewClientV2(endpoint, tos.WithRegion(region),
		tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)))

	if err != nil {
		log.Panic(err)
	}
}
