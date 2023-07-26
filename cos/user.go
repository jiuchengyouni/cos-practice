package cos

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"mime/multipart"
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

func (cos *UserCos) Upload(file *multipart.FileHeader, newSavePath string) (err error, saveFilePath string) {
	saveFileName := file.Filename
	saveFileName = time.DateTime + saveFileName
	saveFilePath = newSavePath + saveFileName
	// 3.通过文件流上传对象
	fd, err := file.Open()
	if err != nil {
	}
	defer fd.Close()
	_, err = cos.Object.Put(context.Background(), saveFilePath, fd, nil)
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