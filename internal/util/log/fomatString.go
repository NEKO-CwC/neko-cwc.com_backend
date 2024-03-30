package utillog

import (
	"log"
	"time"
)

func FormatString(path string, errorFunc string, errorMessage string) {
	log.Printf("[%v] 在 %v 文件中的 %v 函数中出现错误： %v \n", time.Now().Format("2006-01-02 15:04:05"), path, errorFunc, errorMessage)
}
