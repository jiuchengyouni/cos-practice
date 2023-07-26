package router

import (
	"cos_practice/cos"
	"cos_practice/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	//读取部分
	r.POST("/putFileTest", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			//错误抛出
		}
		files := form.File["file"]
		var saveFilePaths []string
		userCos := cos.NewUserCos(context)
		for _, file := range files {
			if r, saveFilePath := userCos.Upload(file, "这是标识"); r == nil {
				saveFilePaths = append(saveFilePaths, saveFilePath)
			}
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
