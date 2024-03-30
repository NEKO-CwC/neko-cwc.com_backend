package routersaccount

import (
	"backend/internal/repositories/userInfo"
	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup() func(c *gin.Context) {
	return func(c *gin.Context) {

		var err error
		var signupRequest SignupRequest

		err = c.ShouldBindJSON(&signupRequest)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "请求格式错误，若此消息发生在网页中，请联系网站管理员",
			})
			return
		}

		email := signupRequest.Email
		password := signupRequest.Password

		accountExist, err := repouserinfo.CheckAccountExist(email)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "数据库错误",
			})
		}

		if !accountExist {
			c.JSON(500, gin.H{
				"message": "账号已经存在，请勿重复注册",
			})
			return
		}

		accountState, _ := repouserinfo.Signup(email, password)
		if !accountState {
			c.JSON(507, gin.H{
				"message": "数据库错误，请过段时间再进行尝试，请联系2840720893@qq.com",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
		return
	}
}
