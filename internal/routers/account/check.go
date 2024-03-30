package routersaccount

import (
	"backend/internal/repositories/userInfo"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckIfAccountExist() func(c *gin.Context) {
	return func(c *gin.Context) {
		accountState, _ := repouserinfo.CheckAccountExist(c.Query("email"))
		if accountState {
			c.JSON(200, gin.H{
				"message": "用户存在",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "用户不存在",
			})
		}
	}
}

func CheckUserInfo() func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误",
			})
		}
		userInfo, err := repouserinfo.CheckUserInfo(userId)
		userInfo.Password = ""
		userInfo.Salt = ""
		if err != nil {
			c.JSON(500, gin.H{
				"message": "数据库出现错误",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "请求成功",
			"data": map[string]any{
				"userInfo": userInfo,
			},
		})
		return
	}
}
