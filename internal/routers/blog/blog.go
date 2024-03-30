package routersblog

import (
	"backend/internal/repositories/blog"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误，请输入正确的文章id",
			})
			return
		}
		blogContent := repoblog.GetBlog(id)
		if blogContent.State == 1 {
			c.JSON(200, gin.H{
				"message": "请求成功",
				"data": map[string]any{
					"blog": blogContent,
				},
			})
			return
		} else {
			c.JSON(500, gin.H{
				"message": "请求失败，此文章不存在",
			})
			return
		}
	}
}
