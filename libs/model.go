// BUG日志分析库文件
// 数据库操作核心代码
// BUG入库的CURD操作

// 数据库字段：
// CREATE TABLE `im_bug` (
//		`bug_id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
//		`title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'BUG标题',
//		`content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'BUG内容',
//		`ctime` int(10) DEFAULT '0' COMMENT '创建时间',
//		`utime` int(10) DEFAULT '0' COMMENT '修改时间',
//		`status` tinyint(1) DEFAULT '0' COMMENT '状态 0未确认 1已确认',
//		PRIMARY KEY (`bug_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

package libs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// bug结构
type Bug struct {
	BugId   int64  `gorm:"primary_key"` //bugID
	Api     string //接口名：api/public/UploadImage
	Title   string //bug标题
	Content string //bug内容
	Num     int    //重复数
	Date    string //bug时间date
	Ctime   int64  //bug产生时间UNIX
	Utime   int64  //bug解决时间
	Status  int8   //是否解决
}

// 连接数据库
func ConnDB(host, username, password, dbname string) (*gorm.DB, error) {
	// 连接数据库
	sqlconn := username + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlconn)
	if err != nil {
		return nil, err
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	// 更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "im_" + defaultTableName
	}
	return db, nil
}
