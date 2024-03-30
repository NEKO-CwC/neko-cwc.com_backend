package routersComment

import (
	"backend/internal/repositories/comment"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetComments() func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误，请输入正确文章id",
			})
			return
		}
		commentSLice, err := repocomment.GetComments(id)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "请求失败，评论部分无法正确加载",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "请求成功",
			"data": map[string]any{
				"comments": commentSLice,
			},
		})
		return
	}
}

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
		voteUpUserIdSlice, err := repocomment.GetVoteInfo(blogId, voteType)
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
