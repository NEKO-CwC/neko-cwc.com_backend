package routersblog

import (
	repoblog "backend/internal/repositories/blog"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetVoteInfo() func(c *gin.Context) {
	return func(c *gin.Context) {
		voteType := c.Param("voteType")
		blogId, err := strconv.Atoi(c.Query("blogId"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误，请输入正确文章id",
			})
			return
		}
		voteUpUserIdSlice, err := repoblog.GetVoteInfo(blogId, voteType)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "数据库出现错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "请求成功",
			"data":    voteUpUserIdSlice,
		})
	}
}

func Vote() func(c *gin.Context) {
	return func(c *gin.Context) {
		voteType := c.Param("voteType")
		operateType := c.Param("operateType")
		userId, err := strconv.Atoi(c.Param("userId"))
		blogId, err := strconv.Atoi(c.Param("blogId"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误",
			})
			return
		}

		err = repoblog.Vote(userId, blogId, voteType, operateType)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "数据库出现错误",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "操作成功",
			"data": map[string]any{
				"userId": userId,
				"blogId": blogId,
			},
		})
		return
	}
}
