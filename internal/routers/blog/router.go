package routersblog

import (
	"github.com/gin-gonic/gin"
)

func LoadBlogRouter(r *gin.Engine) {
	blogGroup := r.Group("/blog") // 创建路由组，并将结果赋值给变量
	{
		blogGroup.GET("/getBlog", GetBlog())
		blogGroup.GET("/getVoteInfo", GetVoteInfo())
		blogGroup.GET("/vote", Vote())
	}
}
