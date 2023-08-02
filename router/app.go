package router

import (
	"bytes"
	"cos_practice/cos"
	"cos_practice/middlewares"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	//读取部分
	r.POST("/putFileTest", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
		}
		files := form.File["file"]
		var saveFilePaths []string
		userCos := cos.NewUserCos(context)
		for _, file := range files {
			//bytes与io.Reader转化，适合于grpc传输
			fd, err := file.Open()
			data, err := ioutil.ReadAll(fd)
			if err != nil {
				// handle error
			}
			fd2 := bytes.NewReader(data)
			err, saveFilePath := userCos.Upload(fd2, "这是标识")
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"paths": "1" + err.Error(),
				})
				return
			}
			saveFilePaths = append(saveFilePaths, saveFilePath)

		}
		context.JSON(http.StatusOK, gin.H{
			"paths": saveFilePaths,
		})
	})

	//下载部分
	r.GET("/getFileTest", func(context *gin.Context) {
		userCos := cos.NewUserCos(context)
		// 读取文件内容
		data, err := userCos.Download("路径")
		if err != nil {
			context.String(http.StatusInternalServerError, "读取文件内容失败")
			return
		}
		var resp *http.Response
		// 设置响应头信息
		context.Header("Content-Type", resp.Header.Get("Content-Type"))
		context.Data(http.StatusOK, resp.Header.Get("Content-Type"), data)
	})
	return r
}
