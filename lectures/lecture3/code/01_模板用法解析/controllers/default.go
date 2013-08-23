package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// 设置 beego 模板文件
	this.TplNames = "index.tpl"

	/* =============== 简单的模板使用 ===================*/

	// 最简单的字段赋值法
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	// 条件判断（只能为 true 或 false，默认为 false）
	this.Data["TrueCond"] = true
	this.Data["FlaseCond"] = false

	// 嵌套输出 - with
	type u struct {
		Name string
		Age  int
		Sex  string
	}
	user := u{
		Name: "Unknown",
		Age:  20,
		Sex:  "Male",
	}

	this.Data["User"] = user

	// 循环输出 - range
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	this.Data["Nums"] = nums
	users := make([]u, 0, 1)
	users = append(users, user)
	this.Data["Users"] = users

	/* =============== 高级的模板使用 ===================*/

	// 通过模板变量传递值
	this.Data["TplVar"] = "hey guys"

	// 使用 beego 内置模板函数进行数据处理
	this.Data["Html"] = "<div>hello beego</div>"

	// pipeline.
	this.Data["Pipe"] = "<div>This string will be escaped through pipeline and template function</div>"

	// 嵌套模板
}
