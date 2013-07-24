package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// 配置管理
	this.Ctx.WriteString("AppName: " + beego.AppConfig.String("appname") +
		"\nRunMode: " + beego.AppConfig.String("runmode"))

	// 默认参数
	this.Ctx.WriteString("\n\nAppName: " + beego.AppName +
		"\nRunMode: " + beego.RunMode)

	// 日志级别
	beego.Trace("Trace test1")
	beego.Info("Info test1")

	beego.SetLevel(beego.LevelInfo)

	beego.Trace("Trace test2")
	beego.Info("Info test2")
}
