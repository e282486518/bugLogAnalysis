// 获取可执行文件的绝对路径

package libs

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取可执行文件的绝对路径
func GetAbsPath() string {
	execpath, _ := os.Executable()
	execpath = filepath.Dir(execpath) // 获得程序路径
	execpath = strings.Replace(execpath, "\\", "/", -1)
	return execpath
}
