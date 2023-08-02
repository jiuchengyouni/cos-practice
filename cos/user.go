package cos

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"io/ioutil"
	"time"
)

// 封装 COS 客户端的结构体和工厂函数。
type UserCos struct {
	*cos.Client
}

func NewUserCos(ctx context.Context) *UserCos {
	if ctx != nil {
		ctx = context.Background()
	}
	return &UserCos{NewCosClient(ctx)}
}

func (cos *UserCos) Upload(file io.Reader, newSavePath string) (err error, saveFilePath string) {
	saveFileName := time.DateTime
	saveFilePath = newSavePath + saveFileName
	fmt.Println(saveFilePath)
	_, err = cos.Object.Put(context.Background(), saveFilePath, file, nil)
	if err == nil {
		//  上传成功,返回资源的相对路径，这里请根据实际返回绝对路径或者相对路径
		return nil, saveFilePath
	}
	return
}

func (cos *UserCos) Download(saveFilePath string) (data []byte, err error) {
	resp, err := cos.Object.Get(context.Background(), saveFilePath, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 读取文件内容
	data, err = ioutil.ReadAll(resp.Body)
	return
}
