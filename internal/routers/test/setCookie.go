package routerstest

import "github.com/gin-gonic/gin"

func setCookie() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.SetCookie("test", "test", 3600, "/", "", false, false)
		c.JSON(200, gin.H{
			"message": "cookie设置成功",
		})
		return
	}
}
