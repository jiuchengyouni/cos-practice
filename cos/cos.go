package cos

import (
	"context"
	"cos_practice/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

// 创建一个 COS的客户端实例。
func NewCosClient(ctx context.Context) *cos.Client {
	u, _ := url.Parse(config.CosUrl)
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.CosSecretId,
			SecretKey: config.CosSecretKey,
		},
	})
}
