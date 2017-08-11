// BUG日志分析库文件
// 邮件操作核心代码
// 发邮件给多人

package libs

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

// 发送邮件
func SendToMail(to []string, content string) {
	//var conf Conf
	config := GetConfig()

	port, _ := strconv.Atoi(config.Smtp["port"]) //string转化为int
	d := gomail.NewDialer(config.Smtp["host"], port, config.Smtp["username"], config.Smtp["password"])

	m := gomail.NewMessage()
	m.SetHeader("From", config.Smtp["username"])
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "APP BUG logs!")
	m.SetBody("text/html", content)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
