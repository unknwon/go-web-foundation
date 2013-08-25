package main

import (
	"beeblog/controllers"
	"beeblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// ORM 命令行模式，用于自动建表，一般只用一次
	orm.RunCommand()

	// 开启 ORM 调试模式
	orm.Debug = true

	// 注册 beego 路由
	beego.Router("/", &controllers.HomeController{})

	// 启动 beego
	beego.Run()
}
