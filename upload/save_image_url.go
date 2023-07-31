package upload

import (
	"context"
	"fmt"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"github.com/yushengguo557/twitter-space/global"
	"github.com/yushengguo557/twitter-space/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

// SaveImageUrl 保存图片 通过url 将twitter用户头像 保存到火山
func SaveImageUrl(business, imaUrl string) (string, error) {
	rawUrl := fmt.Sprintf("http://%s:%d", global.App.Config.Proxy.Host, global.App.Config.Proxy.Port)
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
	request, err := http.NewRequest("GET", imaUrl, nil)
	if err != nil {
		return "", fmt.Errorf("new request, err: %w", err)
	}
	resp, err := httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("do request, err: %w", err)
	}
	// 1.由图片URL获取图片数据
	//resp, err := http.Get(url)
	//if err != nil {
	//	return "", err
	//}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// 2.图片文件名
	filename := filepath.Base(imaUrl)
	key := business + "/" + utils.GenFilename(filename)

	// 3.上传火山TOS
	input := &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{Bucket: global.App.Config.Tos.BucketName, Key: key},
		Content:             resp.Body,
	}
	_, err = global.App.Tos.PutObjectV2(context.Background(), input)
	if err != nil {
		err = fmt.Errorf("upload to tos, err: %w", err)
		return "", err
	}
	return "https://nft-panda.tos-cn-beijing.volces.com/" + key, err
}
