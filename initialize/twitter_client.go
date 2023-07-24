package initialize

import (
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/yushengguo557/twitter-space/global"
	"github.com/yushengguo557/twitter-space/tw"
	"log"
	"net/http"
	"net/url"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func InitTwitterClient() {
	token := global.App.Config.Twitter.BearerToken
	host := global.App.Config.Twitter.APIHost
	// 设置代理 服务器的地址和端口
	rawUrl := fmt.Sprintf("http://%s:%d", global.App.Config.Proxy.Host, global.App.Config.Proxy.Port)
	log.Println("proxy: ", rawUrl)
	proxyUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	global.App.TwitterClient = &tw.TwitterClient{
		Client: &twitter.Client{
			Authorizer: authorize{
				Token: token,
			},
			Client: &http.Client{
				Transport: transport, // 创建一个使用系统代理的 HTTP 客户端
			},
			Host: host,
		},
	}
}
