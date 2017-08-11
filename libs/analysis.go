// BUG日志分析库文件
// 日志分析引擎核心程序，分析并记录BUG
// 使用正则匹配BUG关键词，并提取BUG内容/BUG数据汇总等

package libs

import (
	"bufio"
	"io"
	"strings"
)

// 匹配关键词
func LogAnalysis(str string, keyword []string) (bool, string) {
	for _, v := range keyword {
		if strings.Contains(str, v) {
			return true, v
		}
	}
	return false, ""
}

// 分析一个日志文件
func Analysis(file string, keywords []string) (map[string]int, error) {
	//文件操作
	var text []byte
	// 打开文件
	fs, err := Open(file)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs) // 读文件缓冲区

	ret := make(map[string]int)
	//fmt.Print(ret)
	for io.EOF != err {
		text, _, err = buf.ReadLine() // 读一行
		boo, kwd := LogAnalysis(string(text), keywords)
		if boo {
			ret[kwd]++
		}
	}
	if len(ret) == 0 {
		return nil, nil
	}
	return ret, nil
}
