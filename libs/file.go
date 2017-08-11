// BUG日志分析库文件
// 文件操作核心代码
// 日志文件的读取操作

package libs

import (
	"os"
)

// 打开文件
func Open(path string) (*os.File, error) {
	fs, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return fs, nil
}
