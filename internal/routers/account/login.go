package routersaccount

import (
	"backend/internal/config"
	repouserinfo "backend/internal/repositories/userInfo"
	"backend/internal/util/preProcess"
	utilverify "backend/internal/util/verify"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		var loginRequest LoginRequest

		err = c.ShouldBindJSON(&loginRequest)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误，若此消息发生在前端，请联系网站管理员",
			})
			return
		}

		email := loginRequest.Email
		password := loginRequest.Password
		userInfo, _, message := repouserinfo.Login(email, password)
		userInfo = preProcess.BasicInfo(userInfo)
		if message == "账号不存在" {
			c.JSON(400, gin.H{
				"message": "账号不存在",
			})
			return
		}
		if message == "密码错误" {
			c.JSON(400, gin.H{
				"message": "密码错误",
			})
			return
		}
		if userInfo.State == 0 {
			c.JSON(507, gin.H{
				"message": "数据库出现错误，请联系管理员稍后再试",
			})
			return
		}

		token := utilverify.GenerateLoginToken(userInfo)
		if token == "" {
			c.JSON(507, gin.H{
				"message": "token生成出现错误，请联系管理员稍后再试",
			})
			return
		}

		c.SetCookie("token_login", token, config.EXPIRETIME, "/", "", false, false)
		c.JSON(200, gin.H{
			"message": "登录成功",
			"data": map[string]any{
				"userInfo": userInfo,
				"token":    token,
			},
		})
		return
	}
}
