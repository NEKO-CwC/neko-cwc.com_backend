package middleverify

import (
	utilverify "backend/internal/util/verify"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	token, err := c.Cookie("token_login")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "未登录，无权访问",
		})
		c.Abort()
		return
	}

	tokenClaims, err := utilverify.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "token无效，请重新登录",
		})
		c.Abort()
		return
	}

	if !time.Now().Before(time.Unix(tokenClaims.StandardClaims.ExpiresAt, 0)) {
		c.JSON(400, gin.H{
			"message": "token过期，请重新登录",
		})
		c.Abort()
		return
	}

	c.Next()
}
