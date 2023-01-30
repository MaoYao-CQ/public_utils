package main

import (
	beego "github.com/beego/beego/v2/server/web"
	godump "github.com/favframework/debug"
	"public_utils/config"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	if config.Debug {
		config.Website.Title = "DEBUG-工具喵"
	}

	godump.Dump(config.Website)

	c.Data["website"] = config.Website
	c.Data["timestamp"] = time.Now().Unix()

	c.TplName = "index.tpl"
}

func (c *HomeController) Google() {
	c.Ctx.WriteString("google-site-verification: google965a213a981ca817.html")



//004
}