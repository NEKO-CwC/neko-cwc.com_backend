package routerstest

import "github.com/gin-gonic/gin"

func LoadTestRouter(g *gin.Engine) {
	testGroup := g.Group("/test")
	{
		testGroup.GET("/setCookie", setCookie())
	}
}
