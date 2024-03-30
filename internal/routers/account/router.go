package routersaccount

import (
	"github.com/gin-gonic/gin"
)

func LoadAccountRouter(r *gin.Engine) {
	group := r.Group("/account")
	{
		group.POST("/login", Login())
		group.POST("/signup", Signup())
		group.GET("/check", CheckIfAccountExist())
		group.GET("/info", CheckUserInfo())
	}
}
