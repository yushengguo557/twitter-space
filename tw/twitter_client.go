package tw

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/g8rswimmer/go-twitter/v2"
)

type TwitterClient struct {
	*twitter.Client
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func NewTwitterClient() *TwitterClient {
	token := "AAAAAAAAAAAAAAAAAAAAANihowEAAAAA35C1dWUSYKgQXrycgSzeXLEhefA%3DP2zDy1iZHWuEvVF2pIZhjHoILRdoEVqtTavEVIwpYzzPGftGVj"

	// 设置代理 服务器的地址和端口
	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	return &TwitterClient{
		&twitter.Client{
			Authorizer: authorize{
				Token: token,
			},
			Client: &http.Client{
				Transport: transport, // 创建一个使用系统代理的 HTTP 客户端
			},
			Host: "https://api.twitter.com",
		},
	}
}
