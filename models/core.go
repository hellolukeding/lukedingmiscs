package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 	gorm支持多种数据库，这里主要介绍mysql,连接mysql主要有两个步骤:

	// 配置DSN (Data Source Name)
	// 使用gorm.Open连接数据库
	// DSN格式如下:
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//e.g
	//dsn := "root:123456@tcp(192.168.0.6:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	username := "root"      //账号
	password := "root"      //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "gin_test_db" //数据库名
	timeout := "10s"        //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	//连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

}
