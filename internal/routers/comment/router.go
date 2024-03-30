package routersComment

import (
	modelcomment "backend/internal/models/comment"
	"backend/internal/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoadCommentRouter(r *gin.Engine) {
	commentGroup := r.Group("/comment")
	{
		commentGroup.GET("/getComment", GetComments())
		commentGroup.GET("/getVoteInfo", GetVoteInfo())
		commentGroup.GET("/vote", Vote())
		commentGroup.GET("/reply", Reply())
		commentGroup.GET("/fastReply", func(c *gin.Context) {
			err := repositories.CommentDB.Model(&modelcomment.Comment{}).Create(&modelcomment.Comment{
				ReferId:       -1,
				BlogId:        1,
				UserId:        1,
				UserName:      "NEKO",
				Identity:      "站长",
				Content:       "测试评论",
				CreatedAt:     "114-51-41 91:91:81",
				VoteUpCount:   0,
				VoteDownCount: 0,
				State:         1,
			}).Error

			if err != nil {
				c.JSON(500, gin.H{
					"message": "错误",
				})
				fmt.Println(err.Error())
				return
			}
			c.JSON(200, gin.H{
				"message": "chenggong",
			})
		})
	}
}
