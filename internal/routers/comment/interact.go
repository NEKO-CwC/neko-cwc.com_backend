package routersComment

import (
	modeluser "backend/internal/models/user"
	repocomment "backend/internal/repositories/comment"
	"github.com/gin-gonic/gin"
	"strconv"
)

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

		err = repocomment.Vote(userId, blogId, voteType, operateType)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "数据库发生错误，请联系管理员",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "操作成功",
			"data": map[string]any{
				"userId":      userId,
				"blogId":      blogId,
				"voteType":    voteType,
				"operateType": operateType,
			},
		})
		return
	}
}

type replyStruct struct {
	UserInfo modeluser.Info `json:"userInfo"`
	BlogId   int            `json:"blogId"`
	ReferId  int            `json:"referId"`
	Content  string         `json:"content"`
}

func Reply() func(c *gin.Context) {
	return func(c *gin.Context) {
		var body replyStruct
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误",
			})
			return
		}
		err = repocomment.Reply(body.UserInfo, body.BlogId, body.ReferId, body.Content)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "数据库错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "评论成功",
			"data": map[string]string{
				"user":    body.UserInfo.Name,
				"comment": body.Content,
			},
		})
		return
	}
}
