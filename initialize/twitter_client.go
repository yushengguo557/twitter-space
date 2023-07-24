package initialize

import (
	"context"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/yushengguo557/twitter-space/global"
	"github.com/yushengguo557/twitter-space/tw"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func InitTwitterClient() {
	// 1.获取配置信息
	token := global.App.Config.Twitter.BearerToken
	host := global.App.Config.Twitter.APIHost

	// 2.设置代理 服务器的地址和端口
	rawUrl := fmt.Sprintf("http://%s:%d", global.App.Config.Proxy.Host, global.App.Config.Proxy.Port)
	log.Println("proxy: ", rawUrl)
	proxyUrl, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatal(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	httpClient := &http.Client{
		Transport: transport, // 创建一个使用系统代理的 HTTP 客户端
	}

	// 3.测试代理
	var req *http.Request
	req, err = http.NewRequest("GET", "https://twitter.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	// 使用 context 包创建一个带有超时时间的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	// 4.赋值给全局变量
	global.App.TwitterClient = &tw.TwitterClient{
		Client: &twitter.Client{
			Authorizer: authorize{
				Token: token,
			},
			Client: httpClient,
			Host:   host,
		},
	}
}
