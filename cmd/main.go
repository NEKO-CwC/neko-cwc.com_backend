package main

import (
	routersaccount "backend/internal/routers/account"
	routersblog "backend/internal/routers/blog"
	routerscomment "backend/internal/routers/comment"
	routersfile "backend/internal/routers/file"
	routerstest "backend/internal/routers/test"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {

	// 初始化全局变量

	// 初始化日志文件状态
	file, err := os.OpenFile("main.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("无法打开日志文件" + err.Error())
		return
	}
	defer file.Close()
	log.SetOutput(file)

	// 初始化路由挂载行为
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" || origin == "http://localhost:3001"
		},
		MaxAge: 12 * time.Hour,
	}))
	routersblog.LoadBlogRouter(r)
	routersaccount.LoadAccountRouter(r)
	routerscomment.LoadCommentRouter(r)
	routersfile.LoadFileRouter(r)
	routerstest.LoadTestRouter(r)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
