package main

import (
	"bugLogAnalysis/libs"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 主程序
func main() {
	// 路由设置
	http.HandleFunc("/list", bugList)
	http.HandleFunc("/submit", submit)

	// 设置监听的端口
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// BUG列表
func bugList(w http.ResponseWriter, r *http.Request) {
	// 读配置文件
	config := libs.GetConfig()

	// 连接数据库
	db, err := libs.ConnDB(config.Db["host"], config.Db["username"], config.Db["password"], config.Db["dbname"])
	if err != nil {
		fmt.Print("连接数据库失败")
	}
	defer db.Close()

	// 设置日期
	r.ParseForm()
	date := r.Form.Get("date")
	if date == "" {
		date = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	}

	// 获取数据库数据
	var data []libs.Bug
	db.Where("date = ?", date).Find(&data)

	//fmt.Printf("%v \n", data)

	// 渲染html
	renderHTML(w, "index.html", data)
}

// ajax提交修改状态信息
func submit(w http.ResponseWriter, r *http.Request) {
	// 获取GET参数
	r.ParseForm()
	bugid, err := strconv.Atoi(r.Form.Get("bugid"))
	status, err := strconv.Atoi(r.Form.Get("status"))
	if status == 0 {
		status = 1
	} else {
		status = 0
	}

	// 读配置文件
	config := libs.GetConfig()

	// 连接数据库
	db, err := libs.ConnDB(config.Db["host"], config.Db["username"], config.Db["password"], config.Db["dbname"])
	if err != nil {
		fmt.Print("连接数据库失败")
	}
	defer db.Close()

	// 修改数据库字段值
	var data libs.Bug
	db.Find(&data, bugid)
	err = db.Model(&data).Updates(map[string]interface{}{"utime": time.Now().Unix(), "status": status}).Error

	// 返回json数据
	jsons := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": status,
	}

	if err != nil {
		jsons["code"] = 1
	}

	jstr, _ := json.Marshal(jsons)

	fmt.Fprintf(w, string(jstr))

}

// 渲染页面并输出
func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	tmpl, err := template.ParseFiles("./templates/" + file)
	checkErr(err)
	// 将页面渲染后反馈给客户端
	tmpl.Execute(w, data)
}

// 错误记录
func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
