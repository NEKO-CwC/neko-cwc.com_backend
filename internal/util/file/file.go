package file

import "os"

func Exist(path string) bool {
	_, err := os.Stat(path)
	// 如果os.Stat返回的错误是因为文件或目录不存在，则返回false
	if os.IsNotExist(err) {
		return false
	}
	// 如果没有错误，表示文件或目录存在，或者因为其他原因无法确定，也默认认为存在
	return true
}
