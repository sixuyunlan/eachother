package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist/*
var FS embed.FS //设置静态文件目录

func main() {
	go func() { //新开一个携程，用于监听
		gin.SetMode(gin.DebugMode)                       //设置gin模式为debug
		router := gin.Default()                          //创建一个gin路由
		staticFiles, _ := fs.Sub(FS, "frontend/dist")    //获取静态文件目录
		router.StaticFS("/static", http.FS(staticFiles)) //设置静态文件目录

		router.NoRoute(func(c *gin.Context) { //设置404页面
			path := c.Request.URL.Path        //获取请求路径
			if strings.HasPrefix(path, "/") { //如果请求路径是静态文件目录
				reader, _ := staticFiles.Open("index.html")                            //打开静态文件目录下的index.html
				defer reader.Close()                                                   //关闭文件
				stat, _ := reader.Stat()                                               //获取文件信息
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil) //设置响应内容
			} else { //如果请求路径不是静态文件目录
				c.Status(http.StatusNotFound) //设置响应状态码为404
			}
		}) //设置404页面

		router.Run(":8080") //启动服务
	}() //新开一个携程，用于监听
	chSignal := make(chan os.Signal, 1)                      //创建一个信号接收通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM) //注册信号接收函数

	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"        //设置Chrome路径
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/static/index1.html") //创建Chrome命令
	cmd.Start()                                                                       //启动Chrome

	select { //等待信号
	case <-chSignal: //如果收到信号

	}
	cmd.Process.Kill() //关闭Chrome

}
