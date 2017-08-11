package main

import (
	"bugLogAnalysis/libs"
	"fmt"
	"strconv"
	"time"
)

// 全局变量
var (
	bug_num int = 0 // 当日总bug数
)

// 主程序
func main() {
	defftime := time.Now().UnixNano()

	// 读配置文件
	config := libs.GetConfig()
	//fmt.Printf("%v\n\n", config)

	// 连接数据库
	db, err := libs.ConnDB(config.Db["host"], config.Db["username"], config.Db["password"], config.Db["dbname"])
	if err != nil {
		fmt.Print("连接数据库失败")
	}
	defer db.Close()

	date := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fmt.Printf("处理的日期为：%v \n", date)

	for _, filelogs := range config.Logfile {
		api := filelogs.Api
		keywords := filelogs.Keyword
		maps, _ := libs.Analysis(filelogs.Path+"/"+date+".log", keywords)
		if maps != nil {
			fmt.Printf("%v => %v \n", api, maps)
			// 入库
			for keyword, num := range maps {
				bug_num++ //当日总bug数+1
				logs := libs.Bug{
					Api:     api,
					Title:   keyword,
					Content: "接口：" + api + ", 产生了【" + keyword + "】BUG，共" + strconv.Itoa(num) + "次",
					Num:     num,
					Date:    date,
					Ctime:   time.Now().Unix(),
					Utime:   0,
					Status:  0,
				}
				db.Create(&logs)
			}
		}
	}

	// 发送邮件,耗时2秒多
	libs.SendToMail(config.Mailto, "<h1>"+date+" BUG数汇总</h1><div>今日总bug数有"+strconv.Itoa(bug_num)+"个，请在 http://bugs.xxxxx.com/list?date="+date+" 中查看。</div>")

	// 显示程序执行效率
	defftime = (time.Now().UnixNano() - defftime) / 1e6
	fmt.Printf("程序共执行 %v ms \n", defftime)

}
