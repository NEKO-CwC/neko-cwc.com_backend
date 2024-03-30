package routersfile

import "github.com/gin-gonic/gin"

func LoadFileRouter(r *gin.Engine) {
	fileGroup := r.Group("/file")
	{
		fileGroup.GET("/avatar", GetAvatar())
	}
}
