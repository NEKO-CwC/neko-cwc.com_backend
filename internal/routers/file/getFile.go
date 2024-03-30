package routersfile

import (
	"backend/internal/util/file"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GetAvatar() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Query("id")
		_, filename, _, _ := runtime.Caller(0)
		root := path.Dir(path.Dir(path.Dir(path.Dir(filename))))
		avatarPath := filepath.Join(root, "./fileSource/img/avatar/"+id+".png")
		if file.Exist(avatarPath) {
			imageData, _ := os.ReadFile(avatarPath)
			c.Data(http.StatusOK, "image/png", imageData)
			return
		} else {
			c.AbortWithStatus(http.StatusNotExtended)
			return
		}
	}
}
