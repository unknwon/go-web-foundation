package controllers

import (
	"path"
	"strings"

	"github.com/astaxie/beego"

	"beeblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	lable := this.Input().Get("lable")

	// 获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, category, lable, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, lable, content, attachment)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplNames = "topic_add.html"
	this.Data["IsLogin"] = true
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplNames = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.Data["IsLogin"] = true
}

func (this *TopicController) View() {
	this.TplNames = "topic_view.html"

	reqUrl := this.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Lables"] = strings.Split(topic.Lables, " ")

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}
